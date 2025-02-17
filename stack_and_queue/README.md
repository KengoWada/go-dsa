# Stacks and Queues

## Stacks

A stack is a linear data structure that follows the LIFO (Last In, First Out)
principle, meaning the last element added is the first one to be removed.

Although a stack is a fundamental data structure, it is considered abstract
because it can be implemented using arrays or linked lists.

A stack supports three primary operations:

1. Push – Adds an element to the top of the stack.
1. Pop – Removes and returns the top element from the stack.
1. Peek (or Top) – Returns the top element without removing it.

Since elements are always added and removed from the top, stacks are widely used in function calls, undo mechanisms, and expression evaluation.

## Queues

```sh
                    Peek
                     ↑
Enqueue -> | 1 | 2 | 3 | -> Dequeue

```

A queue is a linear data structure that follows the FIFO (First In, First Out) principle,
meaning the first element added is the first one to be removed.

Queues can be fixed (with a predefined size) or dynamic (growing as needed).
They are commonly implemented using arrays or linked lists.

A queue supports three primary operations:

1. Enqueue – Adds an element to the rear of the queue.
1. Dequeue – Removes and returns the front element of the queue.
1. Peek (or Front) – Returns the front element without removing it.

Queues are widely used in scheduling tasks, handling requests in servers, and breadth-first search (BFS) algorithms.