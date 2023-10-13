> [Problem
> statement](https://en.wikipedia.org/wiki/Dining_philosophers_problem#Problem_statement) 
> 
> Five philosophers dine together at the same table. Each philosopher has their own plate at the table. There is a fork between each plate. 
> The dish served is a kind of spaghetti which has to be eaten with two forks. Each philosopher can only alternately think and eat. 
> Moreover, a philosopher can only eat their spaghetti when they have both a left and right fork. Thus two forks will only be available when 
> their two nearest neighbors are thinking, not eating. After an individual philosopher finishes eating, they will put down both forks. 
> The problem is how to design a regimen (a concurrent algorithm) such that no philosopher will starve; i.e., each can forever continue to alternate 
> between eating and thinking, assuming that no philosopher can know when others may want to eat or think (an issue of incomplete information).

> The problem was designed to illustrate the challenges of avoiding deadlock, a system state in which no progress is possible. To see that a proper solution to this 
> problem is not obvious, consider a proposal in which each philosopher is instructed to behave as follows:
>
* think unless the left fork is available; when it is, pick it up;
* think unless the right fork is available; when it is, pick it up;
* when both forks are held, eat for a fixed amount of time;
* put the left fork down;
* put the right fork down;
* repeat from the beginning.

> However, they each will think for an undetermined amount of time and may end up holding a left fork thinking, staring at the right side of the plate, 
> unable to eat because there is no right fork, until they starve. Resource starvation, mutual exclusion and livelock are other 
> types of sequence and access problem. 

```bash
$ go build
```

```bash 
(Before food portions)
$ ./dining_philosophers 
New table arrival! P 0
New table arrival! P 3
EATING -- Tasty leftover food residue on utensils, P 3
New table arrival! P 1
THINKING -- Pontificating brownian motion, P 1
THINKING -- Pontificating brownian motion, P 1
THINKING -- Pontificating brownian motion, P 1
THINKING -- Pontificating brownian motion, P 1
THINKING -- Pontificating brownian motion, P 1
New table arrival! P 2
EATING -- Tasty leftover food residue on utensils, P 2
EATING -- Tasty leftover food residue on utensils, P 0
THINKING -- Pontificating brownian motion, P 1
EATING -- Tasty leftover food residue on utensils, P 1
New table arrival! P 4
EATING -- Tasty leftover food residue on utensils, P 4
FINISHED - Dining Philosophers have finished eating their meals
```

```bash 
(Current - With food portions)
$ ./dining_philosophers 
New table arrival! P 0
New table arrival! P 4
THINKING -- Pontificating brownian motion, P 4
THINKING -- Pontificating brownian motion, P 4
THINKING -- Pontificating brownian motion, P 4
THINKING -- Pontificating brownian motion, P 4
THINKING -- Pontificating brownian motion, P 4
THINKING -- Pontificating brownian motion, P 4
THINKING -- Pontificating brownian motion, P 4
THINKING -- Pontificating brownian motion, P 4
THINKING -- Pontificating brownian motion, P 4
New table arrival! P 3
EATING -- Portion (1)-- Tasty leftover food residue on utensils, P 3
EATING -- Portion (2)-- Tasty leftover food residue on utensils, P 3
EATING -- Portion (3)-- Tasty leftover food residue on utensils, P 3
New table arrival! P 1
EATING -- Portion (1)-- Tasty leftover food residue on utensils, P 0
EATING -- Portion (2)-- Tasty leftover food residue on utensils, P 0
EATING -- Portion (3)-- Tasty leftover food residue on utensils, P 0
ThInKiNg -- Pontificating brownian motion, P 1
EATING -- Portion (1)-- Tasty leftover food residue on utensils, P 1
EATING -- Portion (2)-- Tasty leftover food residue on utensils, P 1
EATING -- Portion (3)-- Tasty leftover food residue on utensils, P 1
New table arrival! P 2
EATING -- Portion (1)-- Tasty leftover food residue on utensils, P 2
EATING -- Portion (2)-- Tasty leftover food residue on utensils, P 2
EATING -- Portion (3)-- Tasty leftover food residue on utensils, P 2
ThInKiNg -- Pontificating brownian motion, P 4
EATING -- Portion (1)-- Tasty leftover food residue on utensils, P 4
EATING -- Portion (2)-- Tasty leftover food residue on utensils, P 4
EATING -- Portion (3)-- Tasty leftover food residue on utensils, P 4
FINISHED - Dining Philosophers have finished eating their meals
```
