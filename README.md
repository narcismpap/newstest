# news (2019, 3h submission)
GO Programming Test, building a simple News Parser API with a three hour deadline.

This is a small application showing off some interesting capabilities of Go. In the same little service, we attempt 
to operate a configurable, complete news feed parser, filter and cache with an async update tool.

## Running
For convenience, a dockerfile has been added to this repository. Just run:

    make up

Now you can browse the API output at http://localhost:9000/all

## Architecture
The design decisions taken here allow us to reach true infinite horizontal scaling. All data is persistent only to 
memory. We have a highly efficient in-memory caching system, allowing the feeds to be updated out-of-band.

1. We spin off a go HTTP server, handling incoming REST requests to our API
2. At the same time, we run a parallel goroutine that async refreshes all data sources, with proper data locking procedures

Combined, these two elements make for a wicked fast API.

## News API
There are two endpoints exposed by this service, to be consumed by remote clients.

### API Endpoints

    /all - lists all articles in DESC order of publishing
    /articles/{provider}/{category} - filters articles by provider AND/OR category


### Filtering Notes
When attempting to filter by provider or category, you can use:

    /articles/A/B - for {provider}=A, {category}=B
    /articles/*/B - for any provider, {category}=B
    /articles/A/* - for {provider}=A, any category
    /articles/*/* - for all results

### API Example

The response below shows the simple-to-understand API output. The syntax is clear, compact and easy to digest by mobile clients via REST endpoints

    {
        "Articles": [
            {
                "Title": "How fish and shrimps could be recruited as underwater spies",
                "Body": "Animals have long been used for military purposes, but could marine creatures also act as sensors?",
                "Time": "2019-06-06T23:10:50Z",
                "Media": "http://c.files.bbci.co.uk/889F/production/_107257943_gettyimages-157349709.jpg",
                "URL": "https://www.bbc.co.uk/news/business-48515956",
                "Category": "Technology",
                "Provider": "BBC"
            },
        ],
        "Categories": [
            "UK",
            "Technology"
        ],
        "Providers": [
            "BBC",
            "Reuters"
        ]
    }
    
### Microservice Design
Due to time constraints, a monolithic service was the only one that made sense here. Given enough time however, this 
can easily be split apart in at least three services: async refresh, data cache and API gateway. 

### Todo's
With more time, the following stretch tasks would be resolved:

* [ ] Write unit tests
* [ ] Write integration tests
* [ ] Improve RSS parsing performance (more custom parsers + benchmark)

    
### Final Notes
This was fun, coding never ceases to be a great exercise. XML parsing, on the other end, sucks.
