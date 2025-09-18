import (
	"container/heap"
)

// --- Max Heap ---
type Node struct {
	userId   int
	taskId   int
	priority int
	idx      int
}

type PriorityQueue []*Node

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	if pq[i].priority == pq[j].priority {
		return pq[i].taskId > pq[j].taskId
	}
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].idx = i
	pq[j].idx = j
}

func (pq *PriorityQueue) Push(x any) {
	n := x.(*Node)
	n.idx = len(*pq)

	*pq = append(*pq, n)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq // 값복사?
	n := len(old)
	node := old[n-1]
	old[n-1] = nil
	node.idx = -1
	*pq = old[0 : n-1]
	return node
}

// --- TaskManager ---
type TaskManager struct {
	task map[int]*Node
	pq   *PriorityQueue
}

func Constructor(tasks [][]int) TaskManager {
	pq := &PriorityQueue{}
	heap.Init(pq)

	tm := TaskManager{
		task: make(map[int]*Node, len(tasks)),
		pq:   pq,
	}

	for _, t := range tasks {
		uid, tid, p := t[0], t[1], t[2]

		n := &Node{userId: uid, taskId: tid, priority: p}
		tm.task[tid] = n
		heap.Push(tm.pq, n)
	}
	return tm
}

func (tm *TaskManager) Add(userId int, taskId int, priority int) {
	n := &Node{userId: userId, taskId: taskId, priority: priority}
	tm.task[taskId] = n
	heap.Push(tm.pq, n)
}

func (tm *TaskManager) Edit(taskId int, newPriority int) {
	node := tm.task[taskId]
	node.priority = newPriority
	heap.Fix(tm.pq, node.idx)
}

func (tm *TaskManager) Rmv(taskId int) {
	node := tm.task[taskId]

	if node.idx >= 0 {
		heap.Remove(tm.pq, node.idx)
	}
	delete(tm.task, taskId)
}

func (tm *TaskManager) ExecTop() int {
	for tm.pq.Len() > 0 {
		top := heap.Pop(tm.pq).(*Node)
		return top.userId
	}

	return -1
}