# Description

I would like to save data in a compress form in my mongo database IOT save space.

Some fields are periodically present in my data and I would like to use a shorter form for them.

## Goal of project

Create a mechanism that convert data from **string** to **integer** and save it to our database.

### How it works

In case I want to convert from **string**.

```
> make database interrogation with **my_string**

> in case **my_string** is present, then return **integer_value_of_string**

> in case **my_string** is not present, insert it into my database with the next value
(my collection will have a counter for that will tell us the collection size, next value will
have te counter value)

> make database interrogation with **my_integer**
(usually those kind o interrogation need to be done when I want to decompress my data.)

> in case **my_integer** is present, then return **string_value_of_integer**

> in case **my_integer** is not present, raise an error message
```

## How to use

Todo:
