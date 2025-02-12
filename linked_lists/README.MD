# Linked Lists

A linked list is a linear data structure consisting of a sequence of nodes,
where each node contains two components: a data value and a pointer to the
next node in the sequence.

The linked list starts with a head, which is a reference to the first node
in the list. The last node, known as the tail, contains a pointer set to
null, indicating the end of the list.

Linked lists come in different types, including singly linked lists (where
each node points to the next one), doubly linked lists (where nodes also
maintain a reference to the previous node), and circular linked lists
(where the last node links back to the first).

## Types of Linked List

### Singly Linked List

```sh
head -> |1|-> |2|-> |3|-> null
```

A singly linked list is a type of linked list in which each node contains two
components: a data value and a pointer to the next node in the sequence.
The last node in the list points to null, indicating the end of the structure.

Due to this one-way linking, a singly linked list can only be traversed in a
single direction, from the head to the tail.

### Doubly Linked List

```sh
       |Next|->|Next|->|Next|->null
head|->|1   |  |2   |  |3   |
 null<-|Prev|<-|Prev|<-|Prev|
```

A doubly linked list is a type of linked list in which each node contains three components:
a data value, a pointer to the next node, and a pointer to the previous node.

In a doubly linked list:

- The head node's previous pointer is set to null, indicating the start of the list.

- The tail node's next pointer is set to null, marking the end of the list.

Unlike a singly linked list, a doubly linked list allows traversal in both directions—forward
from the head to the tail or backward from the tail to the head.

### Circular Linked List

A circular linked list is a type of linked list in which the tail node points back to the
head node instead of null, forming a continuous loop.

A circular linked list can be either singly linked or doubly linked:

- In a singly circular linked list, each node has a pointer to the
next node, and the last node points back to the first.

```sh
head -> |1|-> |2|-> |3|--|
         ↑               |
         | ______________|
```

- In a doubly circular linked list, each node has pointers to both the
next and previous nodes, with the head node’s previous pointer connecting
to the tail and the tail node’s next pointer connecting to the head.

```sh
         |--------------------|
         ↓                    |
       |Next|->|Next|->|Next|--
head|->|1   |  |2   |  |3   |
    ---|Prev|<-|Prev|<-|Prev|
    |                    ↑
    |____________________|
```

Due to this circular structure, traversal can become indefinite unless a specific stopping condition is defined.
