March 3
Using bubble tea package to design the terminal UI

Basic UI

--------------------------------------------------------------------------------------------
The text to be typed out by the user ksdjbfksdf kjdfksfkdsnf ....

==============================================================================
> [This space will be where the user will be typing out the above passage]

---------------------------------------------------------------------------------------------

As the user types out the passage, show the typed out content word by word and also highlight the corresponding
word in the actual passage with color coding to show if the user is typing it out correctly or not


How to track words per minute - need to keep track of time in a diff goroutine 
One goroutine to render the typed out word and color code the corresponding word in the passage
One goroutine to keep track of number of words typed, number of correct and incorrect words(can this be done in 1 goroutine?)
Also need to maintain a timer and give an option to end test early


Step 1: Design bubbletea based TUI

- Defining the "state" of the TUI
State will comprise of: 
1. The passage text
2. Current word being typed out(down to the current letter which needs to be typed out)
3. Number of correct and incorrect words prev typed out (as a map ie: {"correct": [0,2,4], "incorrect":[1,3]} where the numbers are the indices of the previously typed out words - instead of map, can use two lists as well)

How state will change
- When some letter is typed out - current word state changes(as current letter changes) - based on this the coloring of current word changes
- When space is pressed, the correct + incorrect word mapping changes
- On end of time or if user preemptively ends test, state changes - we will display a new screen showing the result 

How to render the TUI?

What will trigger an update of the model?
- Time running out/ User deciding to end the test
- User typing some letter - updates currentTypedWord
- User pressing space - updates currentTypedWord, count of correct and incorrect prev words, currentPassageTextIndex

TODO: On every key press which is a handled key(letter or space), do a re-render of the screen. What happens in this re-render:
- If space, what is already happening - we clear out 
TODO: Run the 60s timer and have it send a signal to the main goroutine when time ends - this signal should be in 
the form of a message which is handled in Update