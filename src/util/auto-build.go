package util

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"prts-backend/src/model"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/tidwall/gjson"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var cwd string
var gamedata string
var excel string
var levels string
var data string
var server string

func InitPath() error {
	if cwd != "" &&
		gamedata != "" &&
		excel != "" &&
		levels != "" &&
		data != "" &&
		server != "" {
		return nil
	}

	server = "zh_CN"
	var err error
	cwd, err = os.Getwd()
	if err != nil {
		return err
	}
	gamedata = filepath.Join(cwd, "..", "gamedata", server, "gamedata")
	data = filepath.Join(cwd, "..", "data", server)
	excel = filepath.Join(gamedata, "excel")
	levels = filepath.Join(gamedata, "levels")

	_, err = os.Stat(data)
	if os.IsNotExist(err) {
		if err := os.Mkdir(data, 0666); err != nil {
			return err
		}
	} else {
		return err
	}
	return nil
}

func getPath(dir string, path string) string {
	return filepath.Join(dir, path)
}

func ReadFromExcel(path string) ([]byte, error) {
	return os.ReadFile(getPath(excel, path))
}

func ReadFromLevels(path string) ([]byte, error) {
	return os.ReadFile(getPath(levels, path))
}

func WriteJson(obj interface{}, path string) error {
	bytes, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	err = os.WriteFile(path, bytes, 0666)
	if err != nil {
		return err
	}
	return nil
}

func CreateInBatches(db *gorm.DB, obj interface{}) error {
	return db.
		Clauses(clause.OnConflict{UpdateAll: true}).
		CreateInBatches(obj, 50).
		Error
}

func AutoBuild(db *gorm.DB) error {
	err := InitPath()
	if err != nil {
		return err
	}
	err = BuildEnemy(db)
	if err != nil {
		return err
	}
	err = BuildEnemyInstance(db)
	if err != nil {
		return err
	}
	err = BuildStage(db)
	if err != nil {
		return err
	}
	err = BuildEI_S(db)
	if err != nil {
		return err
	}
	err = BuildItem(db)
	if err != nil {
		return err
	}
	err = BuildDrop(db)
	if err != nil {
		return err
	}
	err = BuildCharacter(db)
	if err != nil {
		return err
	}
	err = BuildCharacterInstance(db)
	if err != nil {
		return err
	}
	err = BuildBuildingSkill(db)
	if err != nil {
		return err
	}
	err = BuildC_BS(db)
	if err != nil {
		return err
	}
	err = BuildSkill(db)
	if err != nil {
		return err
	}
	err = BuildSkillInstance(db)
	if err != nil {
		return err
	}
	err = BuildC_S(db)
	if err != nil {
		return err
	}
	err = BuildTalent(db)
	if err != nil {
		return err
	}
	return nil
}

func BuildEnemy(db *gorm.DB) error {
	if err := InitPath(); err != nil {
		return err
	}

	enemyHandbookTable, err := ReadFromExcel("enemy_handbook_table.json")
	if err != nil {
		return err
	}

	var group []model.Enemy
	gjson.ParseBytes(enemyHandbookTable).
		ForEach(func(_, value gjson.Result) bool {
			var single model.Enemy
			single.Id = value.Get("enemyId").String()
			single.SortId = value.Get("sortId").Int()
			single.Name = value.Get("name").String()
			single.Index = value.Get("enemyIndex").String()
			single.Race = value.Get("enemyRace").String()
			single.Level = value.Get("enemyLevel").String()
			single.AttackType = value.Get("attackType").String()
			single.Endure = value.Get("endure").String()
			single.Attack = value.Get("attack").String()
			single.Defence = value.Get("defence").String()
			single.Resistance = value.Get("resistance").String()
			single.Description = value.Get("description").String()
			single.Ability = value.Get("ability").String()
			single.IsInvalidKilled = value.Get("isInvalidKilled").Bool()
			group = append(group, single)
			return true
		})

	if err := WriteJson(group, getPath(data, "enemy.json")); err != nil {
		return err
	}
	if err := CreateInBatches(db, group); err != nil {
		return err
	}
	return nil
}

func BuildEnemyInstance(db *gorm.DB) error {
	if err := InitPath(); err != nil {
		return err
	}

	enemyDatabase, err := ReadFromLevels(getPath("enemydata", "enemy_database.json"))
	if err != nil {
		return err
	}

	var group []model.EnemyInstance
	gjson.GetBytes(enemyDatabase, "enemies").
		ForEach(func(_, value gjson.Result) bool {
			var single model.EnemyInstance
			single.EnemyId = value.Get("Key").String()
			value.Get("Value").ForEach(func(_, value gjson.Result) bool {
				single.Level = value.Get("level").Int()
				if value.Get("enemyData.name.m_defined").Bool() {
					single.Name = value.Get("enemyData.name.m_value").String()
				}
				if value.Get("enemyData.description.m_defined").Bool() {
					single.Description = value.Get("enemyData.description.m_value").String()
				}
				if value.Get("enemyData.attributes.maxHp.m_defined").Bool() {
					single.MaxHp = value.Get("enemyData.attributes.maxHp.m_value").Int()
				}
				if value.Get("enemyData.attributes.atk.m_defined").Bool() {
					single.Atk = value.Get("enemyData.attributes.atk.m_value").Int()
				}
				if value.Get("enemyData.attributes.def.m_defined").Bool() {
					single.Def = value.Get("enemyData.attributes.def.m_value").Int()
				}
				if value.Get("enemyData.attributes.magicResistance.m_defined").Bool() {
					single.MagicResistance = value.Get("enemyData.attributes.magicResistance.m_value").Float()
				}
				if value.Get("enemyData.attributes.cost.m_defined").Bool() {
					single.Cost = value.Get("enemyData.attributes.cost.m_value").Int()
				}
				if value.Get("enemyData.attributes.blockCnt.m_defined").Bool() {
					single.BlockCnt = value.Get("enemyData.attributes.blockCnt.m_value").Int()
				}
				if value.Get("enemyData.attributes.moveSpeed.m_defined").Bool() {
					single.MoveSpeed = value.Get("enemyData.attributes.moveSpeed.m_value").Float()
				}
				if value.Get("enemyData.attributes.attackSpeed.m_defined").Bool() {
					single.AttackSpeed = value.Get("enemyData.attributes.attackSpeed.m_value").Float()
				}
				if value.Get("enemyData.attributes.baseAttackTime.m_defined").Bool() {
					single.BaseAttackTime = value.Get("enemyData.attributes.baseAttackTime.m_value").Float()
				}
				if value.Get("enemyData.attributes.respawnTime.m_defined").Bool() {
					single.RespawnTime = value.Get("enemyData.attributes.respawnTime.m_value").Float()
				}
				if value.Get("enemyData.attributes.hpRecoveryPerSec.m_defined").Bool() {
					single.HpRecoveryPerSec = value.Get("enemyData.attributes.hpRecoveryPerSec.m_value").Float()
				}
				if value.Get("enemyData.attributes.spRecoveryPerSec.m_defined").Bool() {
					single.SpRecoveryPerSec = value.Get("enemyData.attributes.spRecoveryPerSec.m_value").Float()
				}
				if value.Get("enemyData.attributes.maxDeployCount.m_defined").Bool() {
					single.MaxDeployCount = value.Get("enemyData.attributes.maxDeployCount.m_value").Int()
				}
				if value.Get("enemyData.attributes.massLevel.m_defined").Bool() {
					single.MassLevel = value.Get("enemyData.attributes.massLevel.m_value").Int()
				}
				if value.Get("enemyData.attributes.baseForceLevel.m_defined").Bool() {
					single.BaseForceLevel = value.Get("enemyData.attributes.baseForceLevel.m_value").Int()
				}
				if value.Get("enemyData.attributes.tauntLevel.m_defined").Bool() {
					single.TauntLevel = value.Get("enemyData.attributes.tauntLevel.m_value").Int()
				}
				if value.Get("enemyData.attributes.silenceImmune.m_defined").Bool() {
					single.SilenceImmune = value.Get("enemyData.attributes.silenceImmune.m_value").Bool()
				}
				if value.Get("enemyData.attributes.stunImmune.m_defined").Bool() {
					single.StunImmune = value.Get("enemyData.attributes.stunImmune.m_value").Bool()
				}
				if value.Get("enemyData.attributes.sleepImmune.m_defined").Bool() {
					single.SleepImmune = value.Get("enemyData.attributes.sleepImmune.m_value").Bool()
				}
				if value.Get("enemyData.attributes.frozenImmune.m_defined").Bool() {
					single.FrozenImmune = value.Get("enemyData.attributes.frozenImmune.m_value").Bool()
				}
				if value.Get("enemyData.attributes.levitateImmune.m_defined").Bool() {
					single.LevitateImmune = value.Get("enemyData.attributes.levitateImmune.m_value").Bool()
				}
				if value.Get("enemyData.lifePointReduce.m_defined").Bool() {
					single.LifePointReduce = value.Get("enemyData.lifePointReduce.m_value").Int()
				}
				if value.Get("enemyData.levelType.m_defined").Bool() {
					single.LevelType = value.Get("enemyData.levelType.m_value").Int()
				}
				if value.Get("enemyData.rangeRadius.m_defined").Bool() {
					single.RangeRadius = value.Get("enemyData.rangeRadius.m_value").Float()
				}
				if value.Get("enemyData.numOfExtraDrops.m_defined").Bool() {
					single.NumOfExtraDrops = value.Get("enemyData.numOfExtraDrops.m_value").Int()
				}
				group = append(group, single)
				return true
			})
			return true
		})

	if err := WriteJson(group, getPath(data, "enemy-instance.json")); err != nil {
		return err
	}
	if err := CreateInBatches(db, group); err != nil {
		return err
	}
	return nil
}

func BuildStage(db *gorm.DB) error {
	if err := InitPath(); err != nil {
		return err
	}

	stageTable, err := ReadFromExcel("stage_table.json")
	if err != nil {
		return err
	}
	activityTable, err := ReadFromExcel("activity_table.json")
	if err != nil {
		return err
	}
	zoneToActivity := gjson.GetBytes(activityTable, "zoneToActivity")
	basicInfo := gjson.GetBytes(activityTable, "basicInfo")

	var group []model.Stage
	gjson.GetBytes(stageTable, "stages").
		ForEach(func(_, value gjson.Result) bool {
			var single model.Stage
			single.Id = value.Get("stageId").String()
			single.HardId.String = value.Get("hardStagedId").String()
			single.HardId.Valid = single.HardId.String != ""
			single.Type = value.Get("stageType").String()
			single.Difficulty = value.Get("difficulty").String()
			single.PerformanceStageFlag = value.Get("performanceStageFlag").String()
			single.DiffGroup = value.Get("diffGroup").String()
			single.UnlockCondition = value.Get("unlockCondition").String()
			single.LevelId = value.Get("levelId").String()
			single.ZoneId = value.Get("zoneId").String()
			single.Code = value.Get("code").String()
			single.Name = value.Get("name").String()
			single.Description = value.Get("description").String()
			single.DangerLevel = value.Get("dangerLevel").String()
			single.CanPractice = value.Get("canPractice").Bool()
			single.PracticeTicketCost = value.Get("practiceTicketCost").Int()
			single.ApCost = value.Get("apCost").Int()
			single.Activity = zoneToActivity.Get(single.ZoneId).String()
			single.ActivityName = basicInfo.Get(single.Activity + ".name").String()
			single.ActivityDisplayType = basicInfo.Get(single.Activity + ".displayType").String()
			single.StartTime = time.Unix(basicInfo.Get(single.Activity+".startTime").Int(), 0)
			single.EndTime = time.Unix(basicInfo.Get(single.Activity+".endTime").Int(), 0)
			if single.LevelId != "" {
				LevelPath := filepath.Join(levels, strings.ToLower(single.LevelId)+".json")
				_, err := os.Stat(LevelPath)
				if err != nil && !os.IsNotExist(err) {
					log.Fatal(err)
				} else if err == nil {
					bytes, err := os.ReadFile(LevelPath)
					if err != nil {
						log.Fatal(err)
					}
					value := gjson.GetBytes(bytes, "options")
					single.CharacterLimit = value.Get("characterLimit").Int()
					single.MaxLifePoint = value.Get("maxLifePoint").Int()
					single.InitialCost = value.Get("initialCost").Int()
					single.MaxCost = value.Get("maxCost").Int()
				}
			}
			group = append(group, single)
			return true
		})

	if err := WriteJson(group, getPath(data, "stage.json")); err != nil {
		return err
	}
	if err := CreateInBatches(db, group); err != nil {
		return err
	}
	return nil
}

func BuildEI_S(db *gorm.DB) error {
	if err := InitPath(); err != nil {
		return err
	}

	stageTable, err := ReadFromExcel("stage_table.json")
	if err != nil {
		return err
	}

	var group []model.EI_S
	gjson.GetBytes(stageTable, "stages").
		ForEach(func(_, value gjson.Result) bool {
			stageId := value.Get("stageId").String()
			levelId := value.Get("levelId").String()
			if levelId != "" {
				bytes, err := ReadFromLevels(strings.ToLower(levelId) + ".json")
				if err != nil {
					return true
				}
				gjson.GetBytes(bytes, "enemyDbRefs").
					ForEach(func(_, value gjson.Result) bool {
						var single model.EI_S
						single.EnemyId = value.Get("id").String()
						single.Level = value.Get("level").Int()
						single.StageId = stageId
						group = append(group, single)
						return true
					})
			}
			return true
		})

	if err := WriteJson(group, getPath(data, "ei_s.json")); err != nil {
		return err
	}
	if err := CreateInBatches(db, group); err != nil {
		return err
	}
	return nil
}

func BuildItem(db *gorm.DB) error {
	if err := InitPath(); err != nil {
		return err
	}

	itemTable, err := os.ReadFile(filepath.Join(excel, "item_table.json"))
	if err != nil {
		return err
	}

	var group []model.Item
	gjson.GetBytes(itemTable, "items").
		ForEach(func(_, value gjson.Result) bool {
			var single model.Item
			single.Id = value.Get("itemId").String()
			single.Name = value.Get("name").String()
			single.Description = value.Get("description").String()
			single.Rarity = value.Get("rarity").Int()
			single.IconId = value.Get("iconId").String()
			single.SortId = value.Get("sortId").Int()
			single.Usage = value.Get("usage").String()
			single.ObtainApproach = value.Get("obtainApproach").String()
			single.ClassifyType = value.Get("classifyType").String()
			single.Type = value.Get("itemType").String()
			group = append(group, single)
			return true
		})

	if err := WriteJson(group, getPath(data, "item.json")); err != nil {
		return err
	}
	if err := CreateInBatches(db, group); err != nil {
		return err
	}
	return nil
}

func BuildDrop(db *gorm.DB) error {
	if err := InitPath(); err != nil {
		return err
	}

	resp, err := http.Get("https://penguin-stats.io/PenguinStats/api/v2/result/matrix")
	if err != nil {
		return err
	}
	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = resp.Body.Close()
	if err != nil {
		return err
	}

	var group []model.Drop
	gjson.GetBytes(content, "matrix").
		ForEach(func(_, value gjson.Result) bool {
			var single model.Drop
			single.ItemId = value.Get("itemId").String()
			if single.ItemId == "furni" {
				return true
			}
			single.StageId = value.Get("stageId").String()
			if strings.Contains(single.StageId, "_perm") {
				single.Status = "perm"
				single.StageId = single.StageId[:strings.LastIndex(single.StageId, "_")]
			} else if strings.Contains(single.StageId, "_rep") {
				single.Status = "rep"
				single.StageId = single.StageId[:strings.LastIndex(single.StageId, "_")]
			} else {
				single.Status = "normal"
			}
			ss := value.Get("start").Int()
			single.Start = time.Unix(ss/1000, 0)
			es := value.Get("end").Int()
			single.End.Valid = es != 0
			single.End.Time = time.Unix(es/1000, 0)
			single.Times = value.Get("times").Int()
			single.Quantity = value.Get("quantity").Int()
			single.StdDev = value.Get("stdDev").Float()
			single.UpdateAt = time.Now()
			group = append(group, single)
			return true
		})

	if err := WriteJson(group, getPath(data, "drop.json")); err != nil {
		return err
	}
	if err := CreateInBatches(db, group); err != nil {
		return err
	}
	return nil
}

func BuildCharacter(db *gorm.DB) error {
	if err := InitPath(); err != nil {
		return err
	}

	characterTable, err := ReadFromExcel("character_table.json")
	if err != nil {
		return err
	}
	handbookInfoTable, err := ReadFromExcel("handbook_info_table.json")
	if err != nil {
		return err
	}
	handbookDict := gjson.GetBytes(handbookInfoTable, "handbookDict")

	var group []model.Character
	gjson.ParseBytes(characterTable).
		ForEach(func(key, value gjson.Result) bool {
			var single model.Character
			single.Id = key.String()
			tokenKey := value.Get("tokenKey").String()
			single.TokenId.Valid = tokenKey != ""
			single.TokenId.String = tokenKey
			single.Name = value.Get("name").String()
			single.Description = value.Get("description").String()
			single.CanUseGeneralPotentialItem = value.Get("canUseGeneralPotentialItem").Bool()
			single.CanUseActivityPotentialItem = value.Get("canUseActivityPotentialItem").Bool()
			single.PotentialItemId = value.Get("potentialItemId").String()
			single.ActivityPotentialItemId = value.Get("activityPotentialItemId").String()
			single.NationId = value.Get("nationId").String()
			single.GroupId = value.Get("groupId").String()
			single.TeamId = value.Get("teamId").String()
			single.DisplayNumber = value.Get("displayNumber").String()
			single.Appellation = value.Get("appellation").String()
			single.Position = value.Get("position").String()
			single.TagList = value.Get("tagList").String()
			single.ItemUsage = value.Get("itemUsage").String()
			single.ItemDesc = value.Get("itemDesc").String()
			single.ItemObtainApproach = value.Get("itemObtainApproach").String()
			single.IsNotObtainable = value.Get("isNotObtainable").Bool()
			single.IsSpChar = value.Get("isSpChar").Bool()
			single.MaxPotentialLevel = value.Get("maxPotentialLevel").Int()
			single.Rarity = value.Get("rarity").Int()
			single.Profession = value.Get("profession").String()
			single.SubProfessionId = value.Get("subProfessionId").String()
			single.AllSkillLvlupList = value.Get("allSkillLvlup").String()
			single.PotentialList = value.Get("potentialRanks").String()
			single.Sex = handbookDict.Get(single.Id + ".storyTextAudio.0.stories.0.storyText").String()
			if strings.Contains(single.Sex, "女") {
				single.Sex = "女"
			} else if strings.Contains(single.Sex, "男") {
				single.Sex = "男"
			} else if strings.Contains(single.Sex, "断罪") {
				single.Sex = "断罪"
			}
			group = append(group, single)
			return true
		})

	if err := WriteJson(group, getPath(data, "character.json")); err != nil {
		return err
	}
	if err := CreateInBatches(db, group); err != nil {
		return err
	}
	return nil
}

func BuildCharacterInstance(db *gorm.DB) error {
	if err := InitPath(); err != nil {
		return err
	}

	characterTable, err := ReadFromExcel("character_table.json")
	if err != nil {
		return err
	}

	var group []model.CharacterInstance
	gjson.ParseBytes(characterTable).
		ForEach(func(key, value gjson.Result) bool {
			var single model.CharacterInstance
			single.CharacterId = key.String()
			value.Get("phases").ForEach(func(key, value gjson.Result) bool {
				single.Phase = key.Int()
				single.RangeId = value.Get("rangeId").String()
				single.EvolveCost = value.Get("evolveCost").String()
				value.Get("attributesKeyFrames").ForEach(func(_, value gjson.Result) bool {
					single.Level = value.Get("level").Int()
					data := value.Get("data")
					single.MaxHp = data.Get("maxHp").Int()
					single.Atk = data.Get("atk").Int()
					single.Def = data.Get("def").Int()
					single.MagicResistance = data.Get("magicResistance").Float()
					single.Cost = data.Get("cost").Int()
					single.BlockCnt = data.Get("blockCnt").Int()
					single.MoveSpeed = data.Get("moveSpeed").Float()
					single.AttackSpeed = data.Get("attackSpeed").Float()
					single.BaseAttackTime = data.Get("baseAttackTime").Float()
					single.RespawnTime = data.Get("respawnTime").Int()
					single.HpRecoveryPerSec = data.Get("hpRecoveryPerSec").Float()
					single.SpRecoveryPerSec = data.Get("spRecoveryPerSec").Float()
					single.MaxDeployCount = data.Get("maxDeployCount").Int()
					single.MaxDeckStackCnt = data.Get("maxDeckStackCnt").Int()
					single.TauntLevel = data.Get("tauntLevel").Int()
					single.MassLevel = data.Get("massLevel").Int()
					single.BaseForceLevel = data.Get("baseForceLevel").Int()
					single.StunImmune = data.Get("stunImmune").Bool()
					single.SilenceImmune = data.Get("silenceImmune").Bool()
					single.SleepImmune = data.Get("sleepImmune").Bool()
					single.FrozenImmune = data.Get("frozenImmune").Bool()
					single.LevitateImmune = data.Get("levitateImmune").Bool()
					group = append(group, single)
					return true
				})
				return true
			})
			return true
		})

	if err := WriteJson(group, getPath(data, "character-instance.json")); err != nil {
		return err
	}
	if err := CreateInBatches(db, group); err != nil {
		return err
	}
	return nil
}

func BuildBuildingSkill(db *gorm.DB) error {
	if err := InitPath(); err != nil {
		return err
	}

	buildingData, err := ReadFromExcel("building_data.json")
	if err != nil {
		return err
	}

	var group []model.BuildingSkill
	gjson.GetBytes(buildingData, "buffs").
		ForEach(func(_, value gjson.Result) bool {
			var single model.BuildingSkill
			single.Id = value.Get("buffId").String()
			single.Icon = value.Get("buffName").String()
			single.Name = value.Get("skillIcon").String()
			single.SortId = value.Get("sortId").Int()
			single.Category = value.Get("buffCategory").String()
			single.RoomType = value.Get("roomType").String()
			single.Description = value.Get("description").String()
			group = append(group, single)
			return true
		})

	if err := WriteJson(group, getPath(data, "building-skill.json")); err != nil {
		return err
	}
	if err := CreateInBatches(db, group); err != nil {
		return err
	}
	return nil
}

func BuildC_BS(db *gorm.DB) error {
	if err := InitPath(); err != nil {
		return err
	}

	buildingData, err := ReadFromExcel("building_data.json")
	if err != nil {
		return err
	}

	var group []model.C_BS
	gjson.GetBytes(buildingData, "chars").
		ForEach(func(_, value gjson.Result) bool {
			var single model.C_BS
			single.CharacterId = value.Get("charId").String()
			value.Get("buffChar").ForEach(func(_, value gjson.Result) bool {
				value.Get("buffData").ForEach(func(_, value gjson.Result) bool {
					single.BuildingSkillId = value.Get("buffId").String()
					single.UnlockPhase = value.Get("cond.phase").Int()
					single.UnlockLevel = value.Get("cond.level").Int()
					group = append(group, single)
					return true
				})
				return true
			})
			return true
		})

	if err := WriteJson(group, getPath(data, "c_bs.json")); err != nil {
		return err
	}
	if err := CreateInBatches(db, group); err != nil {
		return err
	}
	return nil
}

func BuildSkill(db *gorm.DB) error {
	if err := InitPath(); err != nil {
		return err
	}

	skillTable, err := ReadFromExcel("skill_table.json")
	if err != nil {
		return err
	}

	var group []model.Skill
	gjson.ParseBytes(skillTable).
		ForEach(func(_, value gjson.Result) bool {
			var single model.Skill
			single.Id = value.Get("skillId").String()
			iconId := value.Get("iconId").String()
			if iconId == "" {
				single.IconId = single.Id
			} else {
				single.IconId = iconId
			}
			group = append(group, single)
			return true
		})

	if err := WriteJson(group, getPath(data, "skill.json")); err != nil {
		return err
	}
	if err := CreateInBatches(db, group); err != nil {
		return err
	}
	return nil
}

func BuildSkillInstance(db *gorm.DB) error {
	if err := InitPath(); err != nil {
		return err
	}

	skillTable, err := ReadFromExcel("skill_table.json")
	if err != nil {
		return err
	}

	var group []model.SkillInstance
	gjson.ParseBytes(skillTable).
		ForEach(func(_, value gjson.Result) bool {
			var single model.SkillInstance
			single.SkillId = value.Get("skillId").String()
			value.Get("levels").ForEach(func(key, value gjson.Result) bool {
				single.Level = key.Int()
				single.Name = value.Get("name").String()
				single.RangeId = value.Get("rangeId").String()
				single.Description = value.Get("description").String()
				single.Type = value.Get("skillType").Int()
				single.DurationType = value.Get("durationType").Int()
				single.SpType = value.Get("spData.spType").Int()
				single.MaxChargeTime = value.Get("spData.maxChargeTime").Int()
				single.SpCost = value.Get("spData.spCost").Int()
				single.InitSp = value.Get("spData.initSp").Int()
				single.Duration = value.Get("duration").Int()
				value.Get("blackboard").ForEach(func(_, value gjson.Result) bool {
					key := strings.ToLower(value.Get("key").String())
					val := strconv.FormatFloat(value.Get("value").Float(), 'f', 2, 64)
					reg := regexp.MustCompile(`{-?` + key + `:?([\d\.]*)(%?)}`)
					if reg != nil {
						result := reg.FindAllStringSubmatch(single.Description, -1)
						if result != nil {
							single.Description = strings.Replace(single.Description, result[0][0], val, -1)
						}
					}
					return true
				})
				group = append(group, single)
				return true
			})
			return true
		})

	if err := WriteJson(group, getPath(data, "skill-instance.json")); err != nil {
		return err
	}
	if err := CreateInBatches(db, group); err != nil {
		return err
	}
	return nil
}

func BuildC_S(db *gorm.DB) error {
	if err := InitPath(); err != nil {
		return err
	}

	characterTable, err := ReadFromExcel("character_table.json")
	if err != nil {
		return err
	}

	var group []model.C_S
	gjson.ParseBytes(characterTable).
		ForEach(func(key, value gjson.Result) bool {
			var single model.C_S
			single.CharacterId = key.String()
			value.Get("skills").ForEach(func(_, value gjson.Result) bool {
				single.SkillId = value.Get("skillId").String()
				single.LvlupCostCond = value.Get("levelUpCostCond").String()
				group = append(group, single)
				return true
			})
			return true
		})

	if err := WriteJson(group, getPath(data, "c_s.json")); err != nil {
		return err
	}
	if err := CreateInBatches(db, group); err != nil {
		return err
	}
	return nil
}

func BuildTalent(db *gorm.DB) error {
	if err := InitPath(); err != nil {
		return err
	}

	characterTable, err := os.ReadFile(filepath.Join(excel, "character_table.json"))
	if err != nil {
		return err
	}

	var group []model.Talent
	gjson.ParseBytes(characterTable).
		ForEach(func(key, value gjson.Result) bool {
			var single model.Talent
			single.CharacterId = key.String()
			value.Get("talents").ForEach(func(_, value gjson.Result) bool {
				value.Get("candidates").ForEach(func(_, value gjson.Result) bool {
					single.TalentId = value.Get("prefabKey").Int()
					single.Phase = value.Get("unlockCondition.phase").Int()
					single.Level = value.Get("unlockCondition.level").Int()
					single.PotentialRank = value.Get("requiredPotentialRank").Int()
					single.Name = value.Get("name").String()
					single.Description = value.Get("description").String()
					single.RangeId = value.Get("rangeId").String()
					group = append(group, single)
					return true
				})
				return true
			})
			return true
		})

	if err := WriteJson(group, getPath(data, "talent.json")); err != nil {
		return err
	}
	if err := CreateInBatches(db, group); err != nil {
		return err
	}
	return nil
}
