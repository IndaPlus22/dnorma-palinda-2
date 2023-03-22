# Written answers

## Task 1

### bug01
1. There is a deadlock. The program is stuck at the channel send operation.
2. Added a goroutine
3. The goroutine sends the string to the channel and so the main goroutitine can now read from it

### bug02
1. The program ends before the last integer can be printed
2. Added a waitgroup
3. The main goroutine doesn't end until the other goroutine is done.



## Task 2
##### Q:
What happens if you switch the order of the statements `wgp.Wait()` and `close(ch)` in the end of the `main` function?

##### A:
Some of the producer goroutines will still be going while the channel closes. They will try to 
send to the closed channel which results in an error.

##### Q:
What happens if you move the `close(ch)` from the `main` function and instead close the channel in the end of the function `Produce`?

##### A:
When the first goroutine finishes it closes the channel. The other goroutines that aren't done try to send to the closed channel, resulting in an error

##### Q:
What happens if you remove the statement `close(ch)` completely?

##### A:
Doesn't affect the results since the program stops shortly after. Bad practice since the consumers
don't know when the producers are done

##### Q:
What happens if you increase the number of consumers from 2 to 4?

##### A:
The program executes faster since more goroutines work on the same amount of data from the channel.

#### Q: 
Can you be sure that all strings are printed before the program stops?

#### A:
No, since the program might end before a consumer has a chance to print the remaining data.
