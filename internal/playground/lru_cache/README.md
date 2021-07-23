# LRU Cache
The Least Recently Used (LRU) caching scheme removes the least recently used 
frame when the cache is full and a new page is referenced which is not there in
the cache. 


## Data Structures
Two data structures to implement an LRU cache:

1. __Queue__: A doubly linked list. The maximum size of the queue will be equal
to the total number of frames available (cache size). Most recently used pages 
will be near front end and least recently pages will be near rear end.
2. __Hash__: With page number as key and address of the corresponding queue 
node as value.


## Action
When a page is referenced, the required page may be in the memory. If it is in 
the memory, we need to detach the node of the list and bring it to the front of
the queue. If the required page is not in memory, we bring that in memory. In 
simple words, we add a new node ot the front of the queue and update the 
corresponding node address in the hash. If the queue is full, we remove a node 
from the rear of the queue, and add the new node to the front of the queue.


## References:
* [Geeks for Geeks Article](https://www.geeksforgeeks.org/lru-cache-implementation/)


