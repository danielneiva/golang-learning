## Linked list

- Stores pointers to the next element
- List contains size, a pointer to the head and a pointer to the tail
- each element of the list is a node
  - Each node contains data, and at least a pointer
  - doubly linked lists contain a pointer to the next element and a pointer to the previous element
- elements are usually not stored in contiguous locations (i.e. not next to each other)
- O(n) to access position n
- O(1) to add elements to beginning or at the end of the list
  - Doesn't require to shift elements in order to insert an element