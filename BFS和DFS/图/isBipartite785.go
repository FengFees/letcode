package main

/*
存在一个 无向图 ，图中有 n 个节点。其中每个节点都有一个介于 0 到 n - 1 之间的唯一编号。给你一个二维数组 graph ，其中 graph[u] 是一个节点数组，由节点 u 的邻接节点组成。形式上，对于 graph[u] 中的每个 v ，都存在一条位于节点 u 和节点 v 之间的无向边。该无向图同时具有以下属性：
不存在自环（graph[u] 不包含 u）。
不存在平行边（graph[u] 不包含重复值）。
如果 v 在 graph[u] 内，那么 u 也应该在 graph[v] 内（该图是无向图）
这个图可能不是连通图，也就是说两个节点 u 和 v 之间可能不存在一条连通彼此的路径。
二分图 定义：如果能将一个图的节点集合分割成两个独立的子集 A 和 B ，并使图中的每一条边的两个节点一个来自 A 集合，一个来自 B 集合，就将这个图称为 二分图 。

如果图是二分图，返回 true ；否则，返回 false 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/is-graph-bipartite
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/

// 颜色标记法 使用dfs
var (
	RED              = 1
	UNCOLORED, GREEN = 0, 2
	color            []int
	valid            bool
)

func isBipartite(graph [][]int) bool {
	valid = true
	color = make([]int, len(graph))

	for i := 0; i < len(graph) && valid; i++ {
		if color[i] == UNCOLORED {
			dfs(i, RED, graph)
		}
	}
	return valid
}

func dfs(node, c int, graph [][]int) {
	color[node] = c
	neColor := RED
	if c == RED {
		neColor = GREEN
	}

	for _, neighbor := range graph[node] {
		if color[neighbor] == UNCOLORED {
			dfs(neighbor, neColor, graph)
			if !valid {
				return
			}
		} else if color[neighbor] != neColor {
			valid = false
			return
		}
	}
}

// bfs
func bfs(graph [][]int) bool {
	color = make([]int, len(graph))
	for i := 0; i < len(graph); i++ {
		if color[i] == UNCOLORED {
			queue := []int{i}
			color[i] = RED
			for j := 0; j < len(queue); j++ {
				node := queue[i]
				neColor := RED
				if color[node] == RED {
					neColor = GREEN
				}
				for _, neighbor := range graph[node] {
					if color[neighbor] == UNCOLORED {
						queue = append(queue, neighbor)
						color[neighbor] = neColor
					} else if color[neighbor] != neColor {
						return false
					}
				}
			}
		}
	}
	return true
}

func main() {

}
