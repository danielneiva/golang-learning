## Heaps
- Introduced by HeapSort
- Used to implement priority queues
- Can be expressed as a ***complete binary tree***
  - Complete trees means that all the levels of the tree are full, except for the talles leaf, wich can have empty nodes, but are filled from as left as possilbe
#### Min heaps
- The root of any give subtree must be less or equal than its children
- ascending priority
- the smallest element is the first to be removed from the heap

#### Max heaps
- The root of any given subtree must be greater or equal than its children
- descending priority
- the biggest element is the first to be removed from the heap

#### Heapify

- After inserting an element, heapify
- In the heapify process, the inserted element is compared to it's parent, and if i'ts bigger (Max Heap) or smaller (Min Heap), its swapped. The process continues until the condition isn't satisfied anymore
- Uppon removing an element, heapify
- When removing the root (the root will always be the one removed), the last element should replace the root. The comparing and swapping beggins, and the root is swapped with the biggest (Max Heap) or smallest (min heap) of it's children. The process repeat's itself until the element doesn't match the condition
- The complexity of inserting or removing an element is O(h), where h is the height of the tree, which is logn, which means the complexity is ***O(log n)***