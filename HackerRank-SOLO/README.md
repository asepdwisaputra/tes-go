# Jangan Dipedulikan Please!

A climber is trying to reach a flag that is some height above the ground. In the attempt to reach the flag, the climber can make any number of jumps up the rock wall where the flag is mounted. Movements can only be made _up_ the wall, and the climber must end at exactly the height of the flag.

There are _2_ types of jumps:
1.  A jump of height _1_.
2.  A jump of height _bigJump_.

Determine the minimum number of jumps it will take the climber to reach the flag's exact height.

**Example:**
_flagHeight = 8_
_bigJump = 3_

The climber starts at height _0_, takes two jumps of height _bigJump_ and two of height _1_ to reach exactly _8_ units in _4_ jumps.

**Function Description:**
Complete the function _jumps_ in the editor below.

_jumps_ has the following parameter(s):
   _int flagHeight:_  an integer, the flag height
   _int_ _bigJump:_  an integer, the height of the second type of jump

**Returns:**
    _int:_ an integer, the minimum number of jumps necessary

**Constraints:**
-   _1 ≤ bigJump, flagHeight ≤ 109_

## Input Format for Custom Testing
Input from stdin will be processed as follows and passed to the function.

The first line contains an integer _flagHeight_.
The second line contains an integer _bigJump_.

## Sample Case 0
**Sample Input 0:**
```
STDIN     Function
-----     --------
3      →  flagHeight = 3
1      →  bigJump = 1
```

**Sample Output 0:**
```
3
```

**Explanation:**
The climber can only jump _1_ unit or _bigJump_ units. With _bigJump = 1_, the climber can only make _1_-unit jumps. It will take _3_ jumps to reach the flag. Sample Case 1

## Sample Case 1
**Sample Input 1:**
```
STDIN     Function
-----     --------
3      →  flagHeight = 3
2      →  bigJump = 2
```

**Sample Output 1:**
```
2
```

**Explanation:**
The climber will jump _bigJump = 2_ units and then _1_ unit to reach the flag with _2_ jumps. Sample Case 2

## Sample Case 2
**Sample Input 2:**
```
STDIN     Function
-----     --------
3      →  flagHeight = 3
3      →  bigJump = 3
```

**Sample Output 2:**
```
1
```

**Explanation:**
The climber will make _1_ jump _bigJump = 3_ units up the wall to reach the flag.

---


**Finished Code:**
```go
flagHeight := 5
bigJump := 2

tbigJump := flagHeight / bigJump

var total int
   if flagHeight > 0 {
      total = 0
}

if flagHeight < bigJump {
   total = 1
}

if flagHeight%bigJump == 0 {
   total = total + tbigJump
} else {
   sisa := flagHeight % bigJump
   total = sisa + tbigJump
}

fmt.Println(total)
```
