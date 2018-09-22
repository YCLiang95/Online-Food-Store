# Test Plan

> Goal: Implement a simple-to-use and robust testing strategy that will ensure no breaking changes will be pushed to production.

Requirements:
1. Must include test strategy, functional/component test descriptions (goals, expectations)
2. Test metrics including success criteria; tools and methodology to be used, … etc.

## Test Strategy

Within our Git strategy, for every commit that is pushed to the master branch, a build will be triggered by Jenkins that will run all of the unit tests. If any of the tests has failed, the build will fail and notify all of the team members. 

Our test strategy consists of three pillars: **Simplicity**, **Reliability**, **Agility**

1. **Simplicity** - test design needs to be simple enough in order to encourage developers to write unit tests for the functionality they've implemented.
2. **Reliability** - test strategy needs to be implemented with a great sense of importance, it should be working flawlessly as it doesn't exist.
3. **Agility** - each test should be responsibly for testing one thing and one thing only, test cases should be flexibly and easily modifieble. 

### Test Components

#### Component 1

##### Description 1

#### Component 2

##### Description 2

## Test Metrics 

### Success Criteria

### Methology