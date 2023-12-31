---
slug: key-error
title: KeyError Problems
tags: ['python', 'KeyErrors']
date: 2023-11-27
lastmod: ["lastmod", ":git", "date", "publishDate"]
---

During my schooling, cleaning was really only covered regarding spreadsheets. Even in the advanced courses, where python dominates, there is little coverage on what to be on the look out for when using pre-compiled datasets and handling any problems you come across. So, as I run into them, I will be logging them here and other posts as required depending on the error type.

<!--more-->

In my case study covering the cost of food and imports, I ran into the KeyError: <key> does not exist problem. Thankfully, this a relatively quick fix. The 'key' here is to get a list of all of the keys in the dataframe.

```
dataframe.keys()
```
This simple command, will print all of the keys in the data frame. Obvious, right? Well, this will also print any leading or trailing whitespaces that may not have been noticed when going through the dataset. To correct this:
```
dataframe.rename(columns=lambda x: x.strip(), inplace=True)
dataframe.keys()
```
We can strip the whitespace and check to make sure they are, indeed, removed.
Once confirmed, we can move on with our analysis.
