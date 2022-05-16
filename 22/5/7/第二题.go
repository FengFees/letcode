/*
 * Copyright (c) Huawei Technologies Co., Ltd. 2019-2021. All rights reserved.
 * Description: 上机编程认证
 * Note: 缺省代码仅供参考，可自行决定使用、修改或删除
 * 只能import Go标准库
 */
package main

import (
	"sort"
)

type deptDemand struct {
	deptId    int // 部门编号
	employNum int // 招聘目标
	progmThd  int // 机考门槛值
	techThd   int // 技面门槛值
}

type candidate struct {
	peopleId   int // 应聘者编号
	progmGrade int // 机试分数
	techGrade  int // 技面分数
}

type candidates []candidate

func (c candidates) Len() int      { return len(c) }
func (c candidates) Swap(i, j int) { c[i], c[j] = c[j], c[i] }
func (c candidates) Less(i, j int) bool {
	if c[i].techGrade > c[j].techGrade {
		return true
	} else if c[i].techGrade == c[j].techGrade {
		if c[i].progmGrade > c[j].progmGrade {
			return true
		} else if c[i].progmGrade == c[j].progmGrade {
			return c[i].peopleId < c[j].peopleId
		}
	}

	return false
}

// 待实现函数，在此函数中填入答题代码
func getRecruitmentResult(deptDemands []deptDemand, candidateAbilities []candidate) [][]int {
	result := make([][]int, len(deptDemands))
	deptPerson := make(map[int]int)
	employAllNum := 0
	for _, demand := range deptDemands {
		// 招聘数量
		deptPerson[demand.deptId] = demand.employNum
		employAllNum += demand.employNum
	}

	candidateMap := make(map[int]candidate)
	for _, ability := range candidateAbilities {
		// 候选者的map匹配
		candidateMap[ability.peopleId] = ability
	}

	// 招满或者无人可录取时退出
	for employAllNum != 0 {
		// 记录是否无人可录取
		needNotEmploy := true
		// 遍历部门
		for _, demand := range deptDemands {
			// 如果指标>0 则进行选择
			if personNum := deptPerson[demand.deptId]; personNum > 0 {
				var employQueue []candidate
				// 判断是否有符合要求的候选者
				for _, ability := range candidateAbilities {
					if person, ok := candidateMap[ability.peopleId]; ok {
						if person.techGrade >= demand.techThd && person.progmGrade >= demand.progmThd {
							employQueue = append(employQueue, person)
						}
					}
				}

				// 无候选者满足要求，则进行下一个部门招聘
				if len(employQueue) == 0 {
					continue
				}

				needNotEmploy = false
				// 进行候选者队列排序
				sort.Sort(candidates(employQueue))
				// 部门指标减一
				deptPerson[demand.deptId]--
				// 整体指标减一
				employAllNum--
				// 结果记录
				result[demand.deptId] = append(result[demand.deptId], employQueue[0].peopleId)
				// map中移除候选者
				delete(candidateMap, employQueue[0].peopleId)
			}
		}

		// 无人可录取的时候 退出
		if needNotEmploy {
			break
		}
	}

	//fmt.Println(result)

	// 补录
	for _, demand := range deptDemands {
		if len(result[demand.deptId]) > 0 {
			// 如果这个部门有招到人，则取最后一名候选者
			endEmploy := result[demand.deptId][len(result[demand.deptId])-1]
			for _, ability := range candidateAbilities {
				if person, ok := candidateMap[ability.peopleId]; ok {
					// 如果相等，则录取
					if person.techGrade == candidateAbilities[endEmploy].techGrade &&
						person.progmGrade == candidateAbilities[endEmploy].progmGrade {
						result[demand.deptId] = append(result[demand.deptId], person.peopleId)
						// 把这个候选者删除
						delete(candidateMap, person.peopleId)
						// 招一人就结束
						break
					}
				}
			}
		}
	}

	//fmt.Println(result)

	return result
}

//func main() {
//	reader := bufio.NewReader(os.Stdin)
//	deptNum := readInputInt(reader)
//	deptDemands := readDeptDemandsList(reader, deptNum) // 每个部门的要求
//	candidateNum := readInputInt(reader)
//	candidateAbilities := readCandidateList(reader, candidateNum) // 应聘者的信息
//	result := getRecruitmentResult(deptDemands, candidateAbilities)
//	for _, val := range result {
//		fmt.Println(fmt.Sprint(val))
//	}
//}
//
//func readInputInt(reader *bufio.Reader) int {
//	var num int
//	if _, err := fmt.Fscanf(reader, "%d\n", &num); err != nil {
//		fmt.Println(err.Error())
//		return 0
//	}
//	return num
//}
//
//func readDeptDemandsList(reader *bufio.Reader, num int) []deptDemand {
//	deptDemands := make([]deptDemand, num)
//	for i := 0; i < num; i++ {
//		lineBuf, err := reader.ReadString('\n')
//		if err != nil && err != io.EOF {
//			fmt.Println(err.Error())
//			return nil
//		}
//		lineBuf = strings.TrimRight(lineBuf, "\r\n")
//		lineBuf = strings.TrimSpace(lineBuf)
//		intNums := map2IntArray(lineBuf, " ")
//		if len(intNums) != 3 {
//			fmt.Println("input is error")
//			return nil
//		}
//		deptDemands[i] = deptDemand{deptId: i, employNum: intNums[0], progmThd: intNums[1], techThd: intNums[2]}
//	}
//	return deptDemands
//}
//
//func readCandidateList(reader *bufio.Reader, num int) []candidate {
//	candidateAbilities := make([]candidate, num)
//	for i := 0; i < num; i++ {
//		lineBuf, err := reader.ReadString('\n')
//		if err != nil && err != io.EOF {
//			fmt.Println(err.Error())
//			return nil
//		}
//		lineBuf = strings.TrimRight(lineBuf, "\r\n")
//		lineBuf = strings.TrimSpace(lineBuf)
//		intNums := map2IntArray(lineBuf, " ")
//		if len(intNums) != 2 {
//			fmt.Println("input is error")
//			return nil
//		}
//		candidateAbilities[i] = candidate{peopleId: i, progmGrade: intNums[0], techGrade: intNums[1]}
//	}
//	return candidateAbilities
//}
//
//func map2IntArray(str string, dem string) []int {
//	tempArray := strings.Split(str, dem)
//	result := make([]int, len(tempArray))
//	for index, value := range tempArray {
//		value = strings.TrimSpace(value)
//		intVal, err := strconv.Atoi(value)
//		if err == nil {
//			result[index] = intVal
//		}
//	}
//	return result
//}
