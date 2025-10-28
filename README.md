A small (toy) memory store that can hold references to data ordered by tags or as they appear internally "surnames".

These tags are meant to be points from where to potentially grab the data that is related.

Tag would only be viable if the data set is sustainable, it's not meant to be used for very large sets of records.

This works best with small hierarchies that intersect such as:

Given an Entity `User`

```go

type User struct {
    Id string
    Name string
    Email string
    Role string
    // ...
}

```
We could have data structured like this 

- User > Customer > Recurring
- User > Customer > Non-Recurring
- User > Employee > HR > Manager
- User > Employee > Engineering > Backend
- User > Employee > Engineering > Frontend   

Where we can query the full engineering team by simply calling

```go
members := store.Get("engineering")
```

Querying by `User.Id` would as expected yield only the record with that unique id.

 