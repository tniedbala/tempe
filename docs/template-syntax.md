# Template Syntax
Tempe is intended to be a minimal, Jinja-like templating language. Currently this includes 
the following special syntax:

1. [Expressions](#expressions)
2. [Statements](#statements)
    - [Variable Assignment](#variable-assignment)
    - [For Loops](#for-loops)
    - [If Statements](#if-statements)

## Expressions
All text provided within `{{ double curly braces }}` will be evaluated and rendered 
as an expression of the target tempe implementation. For example, this would be any valid starlark 
expression for the [starlark](./../engines/starlark) implementation:

**Template**
```jinja
value = {{ 1 + 2 + 3 }}
```
**Output**
```
value = 6
```

## Statements
Statements are control structures with syntax provided between `{% curly brace + percent signs %}`,
which can be used to assign variables, loop through a collection, or conditionally render 
sub-expressions and/or nested statements.

### Variable Assignment
Variables can be assigned within a template using the following syntax:
```
{% name = <value>; name2 = <value2>; name3 = <value3>; %}
```
- `name` is a variable name, which must begin with a letter or underscore,
followed by zero or more numbers, letters or underscores.
- `<value>` is any valid expression syntax per the target implementation - 
this will include all text lead up to, but not including the following semicolon `;`.
- Every `name = <value>` pair must be followed by a semicolon `;`.
- Any number of variables may be assigned within an assignment statement.
- Any amount of whitespace (including new lines) may be used within an assignment statement.
For example, the following is the same as the above example:
    ```
    {% 
        name = <value>;
        name2 = <value2>; 
        name3 = <value3>; 
    %}
    ```

### For Loops
For loops may be used to generate iterative output, and may contain any amount of nested text, expressions 
and/or statements. Two syntactic variations are available:

1. **Single variable for loop:**
    ```jinja
    {% for name in <values> %}
    ...
    {% endfor %}
    ```
    - `name` is a variable that will be updated with values from the collection being iterated over.
    - `<values>` is any valid iterable expression, per the target implementation.
    - `{% endfor %}` must always be used to close the for loop.

2. **Double-variable for loop:**
    ```jinja
    {% for i, name in <values> %}
    ...
    {% endfor %}
    ```
    - This variation include 2 variable assignments, `i` and `name`. By convention, `i` will typically be the
    index of the current iteration, though `i` and `name` could potentially also represent a key-value pair of 
    a mapping type. This is to be specified by the target implementation.


### If Statements
If statements may be used to conditionally render sections of a template, and may contain any amount of nested text, expressions 
and/or statements. An if statement must contain an `if` clause, which may be followed by zero or more `elseif` clauses, an 
optional `else` clause, and must always be closed using `{% endif %}`.

**Example:**
```jinja
{% if <condition> %}
...
{% elseif <condition2> %}
...
{% elseif <condition3> %}
...
{% else %}
...
{% endif %}
```
- `<condition>` is any valid expression per the target implementation. 

