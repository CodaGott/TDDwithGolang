Writing test in Golang.

This requires a few rules though it's same thing as writing a function in go.

# Rule number 1.
Your Gofile needs to be named xyz_test.go.

# Rule number 2.
The test function must start with the word Test,

#Rule number 3
The test function takes one argument only. `t *testing.T`. 
this is the only param the test func will/should take.

#Rule number 4
In other to use `t *testing.T` type, you need to
import testing in the go file the same way fmt is imported.

After your code passes, you have to now refactor the code as per TDD requirements.

# TDD Discipline
1. Write a test.
2. Make the compiler pass, by not complaining about any forgotten line.
3. Run the test, see that it fails and check the error message is meaningful
4. Write enough code to make the test pass
5. Refactor the code to make it better.