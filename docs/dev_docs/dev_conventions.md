# Preface
*WIP this document will be updated throughout the project*

# Branch & Commit Syntax
For branches we use semantic prefixes (See below).  
[Link to semantic discussion](https://gist.github.com/joshbuchea/6f47e86d2510bce28f8e7f42ae84c716)  
Example:
```
Branches:
    feat_<name_of_feature>
    test_<name_of_test>
    docs_<name_of_docs>
    fix_<name_of_fix>
    style_<name_of_style>
    refactor_<name_of_refactor>
    chore_<name_of_chore>
```
This is the semantic syntax for commits.  
Example:
```
Commits:
    feat: <name_of_feature>
    test: <name_of_test>
    docs: <name_of_docs>
    fix: <name_of_fix>
    style: <name_of_style>
    refactor: <name_of_refactor>
    chore: <name_of_chore> 
```

# Mob & Pair programming conventions
If mob or pair programming is used the driver should be noted in the commit.
Example:
```
feat: Add cool new feature
Driver: <driver_name>

Co-authored-by: <name> <<email>>
...
``` 