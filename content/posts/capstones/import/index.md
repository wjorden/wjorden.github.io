---
slug: google-data-analytics
title: Food Import Cost Case Study
date: 2023-11-12
lastmod: ["lastmod", ":git", "date", "publishDate"]
tags: ["R", "ggplot", "python", "statistics", "matplotlib"]
---

<script src="{{< blogdown/postref >}}index_files/htmlwidgets/htmlwidgets.js"></script>
<script src="{{< blogdown/postref >}}index_files/plotly-binding/plotly.js"></script>
<script src="{{< blogdown/postref >}}index_files/typedarray/typedarray.min.js"></script>
<script src="{{< blogdown/postref >}}index_files/jquery/jquery.min.js"></script>
<link href="{{< blogdown/postref >}}index_files/crosstalk/css/crosstalk.min.css" rel="stylesheet" />
<script src="{{< blogdown/postref >}}index_files/crosstalk/js/crosstalk.min.js"></script>
<link href="{{< blogdown/postref >}}index_files/plotly-htmlwidgets-css/plotly-htmlwidgets.css" rel="stylesheet" />
<script src="{{< blogdown/postref >}}index_files/plotly-main/plotly-latest.min.js"></script>
<script src="{{< blogdown/postref >}}index_files/htmlwidgets/htmlwidgets.js"></script>
<script src="{{< blogdown/postref >}}index_files/plotly-binding/plotly.js"></script>
<script src="{{< blogdown/postref >}}index_files/typedarray/typedarray.min.js"></script>
<script src="{{< blogdown/postref >}}index_files/jquery/jquery.min.js"></script>
<link href="{{< blogdown/postref >}}index_files/crosstalk/css/crosstalk.min.css" rel="stylesheet" />
<script src="{{< blogdown/postref >}}index_files/crosstalk/js/crosstalk.min.js"></script>
<link href="{{< blogdown/postref >}}index_files/plotly-htmlwidgets-css/plotly-htmlwidgets.css" rel="stylesheet" />
<script src="{{< blogdown/postref >}}index_files/plotly-main/plotly-latest.min.js"></script>

This topic was chosen as there is a lot that goes into the production of food. A level of concern was raised as it became obvious that people with various disorders are struggling to get the food items they can have. By this it is meant, those with Crohn’s disease, restrictive diets due to genetics (dysautonomia, EDS, ANS, etc), as well as many others. Medications can make the lack of proper foods tolerable, but it does not solve the overall issue they have.
<!--more-->

# Week 1

## Setup

The libraries in use:

``` r
library(tidyverse)
library(ggplot2)
library(plotly)
library(readxl)
```

## Planning Stage

1)  Find the data set(s) for food production costs (first place to look is USDA)  
2)  Find the data set(s) for consumer cost of the items (USDA?)  
3)  Start plotting the costs against each other (Smooth vs Bar charts)  
4)  Figure out the pattern in the data (likely going to have go back over the course of a few years (on returning, apparently decades))

The main question, why does food cost so much? To answer this question, we need a few bits of data to start with. Those data sets are, the current cost to produce (e.g. grow fruits) and/or manufacture (e.g. process grains) the food items. From there, we need to get a data set for the cost to consumers of these same items. Considering that I live in the United States, I’ll be working with those data sets first.

## Let’s begin

Time to start gathering the data sets. Charts will appear in the order made in the order of the data sets. If more data sets are listed than there are charts, then the initial exploration of the data is not done.

Well, we have two charts made showing the breakdown of imported goods and the year-over-year change in cost. The spikes in import price or the amount of food still needed to be imported raises some concern. Further review will be needed to provide, not only context, but an explanation as to why we saw such a sharp increase.

# End of Week 1 Update

The data is immense. As of right now, the data still requires some restructuring for it to be more accessible for visualization. The Vegetables and Pulses Yearbook, while extensive (50+ tables), is structured for spreadsheet use. Have a happy holiday!

# Week 2

The line chart is the average of imported vegetables and pulses available by US pounds per person. It was made using ggplot2 and a custom data frame.

# End of Week 2

Upon further review, there are a lot of missing pieces. Namely, trade agreements. While we have import costs, this would be offset by trades with respective countries. For example, exporting wheat to import avocados.

# Data sets

[Food Imports](https://www.ers.usda.gov/data-products/u-s-food-imports/)  
[Vegetables and Pulses Yearbook](https://www.ers.usda.gov/data-products/vegetables-and-pulses-data/vegetables-and-pulses-yearbook-tables/)  
[Farm Output](https://www.ers.usda.gov/data-products/agricultural-productivity-in-the-u-s/)  
[Dairy](https://www.ers.usda.gov/data-products/dairy-data.aspx)

# Charts

![2022 US Imports (Spreadsheet)](img/2022-US-Imports.png)  
This is the breakdown of all the imported food and preparatory items, by category.

![US Imported Food YoY (Spreadsheet)](img/US-Imported-Food-Price-Change.png)  
This chart is rather misleading when first looking at it. While the data used to prepare this chart is accurate, it does not take into account trade agreements. There is a possibility that that the true change is much flatter. The data does not make a distinction towards either true expense or expense minus trades.

Now, for some R plots. Sheet 3 is “U.S. per capita availability of selected, commercially produced, fresh and processing vegetables and dry pulse crops”.

``` r
pc <- read_xlsx("datasets/VPPT.xlsx", sheet = 3)

avg <- setNames(aggregate(pc$Value, list(pc$Year), FUN= function(x) round(mean(x), 2)), c("year", "value"))

p <- ggplot(avg, aes(year, value)) +
  geom_line() +
  labs(title = "Average Vegetables and Pulses Per Person") +
  ylab("Weight per Person (lbs)") + xlab("Year")

ggplotly(p)
```

<div class="plotly html-widget html-fill-item" id="htmlwidget-1" style="width:500px;height:400px;"></div>
<script type="application/json" data-for="htmlwidget-1">{"x":{"data":[{"x":[1970,1971,1972,1973,1974,1975,1976,1977,1978,1979,1980,1981,1982,1983,1984,1985,1986,1987,1988,1989,1990,1991,1992,1993,1994,1995,1996,1997,1998,1999,2000,2001,2002,2003,2004,2005,2006,2007,2008,2009,2010,2011,2012,2013,2014,2015,2016,2017,2018,2019,2020,2021,2022],"y":[9.0399999999999991,9.1899999999999995,9.1500000000000004,9.0600000000000005,8.9199999999999999,8.9100000000000001,9.1999999999999993,9.1699999999999999,8.9199999999999999,9.1099999999999994,7.5099999999999998,7.3700000000000001,7.4199999999999999,7.4400000000000004,7.8600000000000003,7.8200000000000003,7.7999999999999998,7.9100000000000001,7.9699999999999998,8.2699999999999996,8.2699999999999996,8.4100000000000001,8.3300000000000001,8.5800000000000001,8.6699999999999999,8.5700000000000003,8.7400000000000002,8.3699999999999992,8.3100000000000005,8.4000000000000004,8.3800000000000008,8.1500000000000004,8.1899999999999995,8.3100000000000005,8.3800000000000008,8.2799999999999994,8.0500000000000007,8.0999999999999996,7.8600000000000003,7.8899999999999997,7.9699999999999998,7.6600000000000001,7.8300000000000001,7.6100000000000003,7.7300000000000004,7.5599999999999996,7.9199999999999999,7.9900000000000002,8.0800000000000001,7.7300000000000004,8.9900000000000002,8.7100000000000009,8.75],"text":["year: 1970<br />value: 9.04","year: 1971<br />value: 9.19","year: 1972<br />value: 9.15","year: 1973<br />value: 9.06","year: 1974<br />value: 8.92","year: 1975<br />value: 8.91","year: 1976<br />value: 9.20","year: 1977<br />value: 9.17","year: 1978<br />value: 8.92","year: 1979<br />value: 9.11","year: 1980<br />value: 7.51","year: 1981<br />value: 7.37","year: 1982<br />value: 7.42","year: 1983<br />value: 7.44","year: 1984<br />value: 7.86","year: 1985<br />value: 7.82","year: 1986<br />value: 7.80","year: 1987<br />value: 7.91","year: 1988<br />value: 7.97","year: 1989<br />value: 8.27","year: 1990<br />value: 8.27","year: 1991<br />value: 8.41","year: 1992<br />value: 8.33","year: 1993<br />value: 8.58","year: 1994<br />value: 8.67","year: 1995<br />value: 8.57","year: 1996<br />value: 8.74","year: 1997<br />value: 8.37","year: 1998<br />value: 8.31","year: 1999<br />value: 8.40","year: 2000<br />value: 8.38","year: 2001<br />value: 8.15","year: 2002<br />value: 8.19","year: 2003<br />value: 8.31","year: 2004<br />value: 8.38","year: 2005<br />value: 8.28","year: 2006<br />value: 8.05","year: 2007<br />value: 8.10","year: 2008<br />value: 7.86","year: 2009<br />value: 7.89","year: 2010<br />value: 7.97","year: 2011<br />value: 7.66","year: 2012<br />value: 7.83","year: 2013<br />value: 7.61","year: 2014<br />value: 7.73","year: 2015<br />value: 7.56","year: 2016<br />value: 7.92","year: 2017<br />value: 7.99","year: 2018<br />value: 8.08","year: 2019<br />value: 7.73","year: 2020<br />value: 8.99","year: 2021<br />value: 8.71","year: 2022<br />value: 8.75"],"type":"scatter","mode":"lines","line":{"width":1.8897637795275593,"color":"rgba(0,0,0,1)","dash":"solid"},"hoveron":"points","showlegend":false,"xaxis":"x","yaxis":"y","hoverinfo":"text","frame":null}],"layout":{"margin":{"t":43.762557077625573,"r":7.3059360730593621,"b":40.182648401826491,"l":43.105022831050235},"plot_bgcolor":"rgba(235,235,235,1)","paper_bgcolor":"rgba(255,255,255,1)","font":{"color":"rgba(0,0,0,1)","family":"","size":14.611872146118724},"title":{"text":"Average Vegetables and Pulses Per Person","font":{"color":"rgba(0,0,0,1)","family":"","size":17.534246575342465},"x":0,"xref":"paper"},"xaxis":{"domain":[0,1],"automargin":true,"type":"linear","autorange":false,"range":[1967.4000000000001,2024.5999999999999],"tickmode":"array","ticktext":["1970","1980","1990","2000","2010","2020"],"tickvals":[1970,1980,1990,2000,2010,2020],"categoryorder":"array","categoryarray":["1970","1980","1990","2000","2010","2020"],"nticks":null,"ticks":"outside","tickcolor":"rgba(51,51,51,1)","ticklen":3.6529680365296811,"tickwidth":0.66417600664176002,"showticklabels":true,"tickfont":{"color":"rgba(77,77,77,1)","family":"","size":11.68949771689498},"tickangle":-0,"showline":false,"linecolor":null,"linewidth":0,"showgrid":true,"gridcolor":"rgba(255,255,255,1)","gridwidth":0.66417600664176002,"zeroline":false,"anchor":"y","title":{"text":"Year","font":{"color":"rgba(0,0,0,1)","family":"","size":14.611872146118724}},"hoverformat":".2f"},"yaxis":{"domain":[0,1],"automargin":true,"type":"linear","autorange":false,"range":[7.2785000000000002,9.2914999999999992],"tickmode":"array","ticktext":["7.5","8.0","8.5","9.0"],"tickvals":[7.5,8,8.5,9],"categoryorder":"array","categoryarray":["7.5","8.0","8.5","9.0"],"nticks":null,"ticks":"outside","tickcolor":"rgba(51,51,51,1)","ticklen":3.6529680365296811,"tickwidth":0.66417600664176002,"showticklabels":true,"tickfont":{"color":"rgba(77,77,77,1)","family":"","size":11.68949771689498},"tickangle":-0,"showline":false,"linecolor":null,"linewidth":0,"showgrid":true,"gridcolor":"rgba(255,255,255,1)","gridwidth":0.66417600664176002,"zeroline":false,"anchor":"x","title":{"text":"Weight per Person (lbs)","font":{"color":"rgba(0,0,0,1)","family":"","size":14.611872146118724}},"hoverformat":".2f"},"shapes":[{"type":"rect","fillcolor":null,"line":{"color":null,"width":0,"linetype":[]},"yref":"paper","xref":"paper","x0":0,"x1":1,"y0":0,"y1":1}],"showlegend":false,"legend":{"bgcolor":"rgba(255,255,255,1)","bordercolor":"transparent","borderwidth":1.8897637795275593,"font":{"color":"rgba(0,0,0,1)","family":"","size":11.68949771689498}},"hovermode":"closest","barmode":"relative"},"config":{"doubleClick":"reset","modeBarButtonsToAdd":["hoverclosest","hovercompare"],"showSendToCloud":false},"source":"A","attrs":{"3fa03b45479b":{"x":{},"y":{},"type":"scatter"}},"cur_data":"3fa03b45479b","visdat":{"3fa03b45479b":["function (y) ","x"]},"highlight":{"on":"plotly_click","persistent":false,"dynamic":false,"selectize":false,"opacityDim":0.20000000000000001,"selected":{"opacity":1},"debounce":0},"shinyEvents":["plotly_hover","plotly_click","plotly_selected","plotly_relayout","plotly_brushed","plotly_brushing","plotly_clickannotation","plotly_doubleclick","plotly_deselect","plotly_afterplot","plotly_sunburstclick"],"base_url":"https://plot.ly"},"evals":[],"jsHooks":[]}</script>

Keep in mind, this is for the <strong>WHOLE YEAR</strong>.

For a month-to-month basis, we would need to divide the value by 12.

``` r
avg_monthly <- setNames(aggregate(pc$Value, list(pc$Year), FUN= function(x) round(mean(x)/12, 2)), c("year", "value"))

p2 <- ggplot(avg_monthly, aes(year, value)) +
  geom_line() +
  labs(title = "Average Vegetables and Pulses Per Person Per Month") +
  ylab("Weight per Person (lbs)") + xlab("Year")

ggplotly(p2)
```

<div class="plotly html-widget html-fill-item" id="htmlwidget-2" style="width:672px;height:480px;"></div>
<script type="application/json" data-for="htmlwidget-2">{"x":{"data":[{"x":[1970,1971,1972,1973,1974,1975,1976,1977,1978,1979,1980,1981,1982,1983,1984,1985,1986,1987,1988,1989,1990,1991,1992,1993,1994,1995,1996,1997,1998,1999,2000,2001,2002,2003,2004,2005,2006,2007,2008,2009,2010,2011,2012,2013,2014,2015,2016,2017,2018,2019,2020,2021,2022],"y":[0.75,0.77000000000000002,0.76000000000000001,0.75,0.73999999999999999,0.73999999999999999,0.77000000000000002,0.76000000000000001,0.73999999999999999,0.76000000000000001,0.63,0.60999999999999999,0.62,0.62,0.65000000000000002,0.65000000000000002,0.65000000000000002,0.66000000000000003,0.66000000000000003,0.68999999999999995,0.68999999999999995,0.69999999999999996,0.68999999999999995,0.71999999999999997,0.71999999999999997,0.70999999999999996,0.72999999999999998,0.69999999999999996,0.68999999999999995,0.69999999999999996,0.69999999999999996,0.68000000000000005,0.68000000000000005,0.68999999999999995,0.69999999999999996,0.68999999999999995,0.67000000000000004,0.67000000000000004,0.66000000000000003,0.66000000000000003,0.66000000000000003,0.64000000000000001,0.65000000000000002,0.63,0.64000000000000001,0.63,0.66000000000000003,0.67000000000000004,0.67000000000000004,0.64000000000000001,0.75,0.72999999999999998,0.72999999999999998],"text":["year: 1970<br />value: 0.75","year: 1971<br />value: 0.77","year: 1972<br />value: 0.76","year: 1973<br />value: 0.75","year: 1974<br />value: 0.74","year: 1975<br />value: 0.74","year: 1976<br />value: 0.77","year: 1977<br />value: 0.76","year: 1978<br />value: 0.74","year: 1979<br />value: 0.76","year: 1980<br />value: 0.63","year: 1981<br />value: 0.61","year: 1982<br />value: 0.62","year: 1983<br />value: 0.62","year: 1984<br />value: 0.65","year: 1985<br />value: 0.65","year: 1986<br />value: 0.65","year: 1987<br />value: 0.66","year: 1988<br />value: 0.66","year: 1989<br />value: 0.69","year: 1990<br />value: 0.69","year: 1991<br />value: 0.70","year: 1992<br />value: 0.69","year: 1993<br />value: 0.72","year: 1994<br />value: 0.72","year: 1995<br />value: 0.71","year: 1996<br />value: 0.73","year: 1997<br />value: 0.70","year: 1998<br />value: 0.69","year: 1999<br />value: 0.70","year: 2000<br />value: 0.70","year: 2001<br />value: 0.68","year: 2002<br />value: 0.68","year: 2003<br />value: 0.69","year: 2004<br />value: 0.70","year: 2005<br />value: 0.69","year: 2006<br />value: 0.67","year: 2007<br />value: 0.67","year: 2008<br />value: 0.66","year: 2009<br />value: 0.66","year: 2010<br />value: 0.66","year: 2011<br />value: 0.64","year: 2012<br />value: 0.65","year: 2013<br />value: 0.63","year: 2014<br />value: 0.64","year: 2015<br />value: 0.63","year: 2016<br />value: 0.66","year: 2017<br />value: 0.67","year: 2018<br />value: 0.67","year: 2019<br />value: 0.64","year: 2020<br />value: 0.75","year: 2021<br />value: 0.73","year: 2022<br />value: 0.73"],"type":"scatter","mode":"lines","line":{"width":1.8897637795275593,"color":"rgba(0,0,0,1)","dash":"solid"},"hoveron":"points","showlegend":false,"xaxis":"x","yaxis":"y","hoverinfo":"text","frame":null}],"layout":{"margin":{"t":43.762557077625573,"r":7.3059360730593621,"b":40.182648401826491,"l":48.949771689497723},"plot_bgcolor":"rgba(235,235,235,1)","paper_bgcolor":"rgba(255,255,255,1)","font":{"color":"rgba(0,0,0,1)","family":"","size":14.611872146118724},"title":{"text":"Average Vegetables and Pulses Per Person Per Month","font":{"color":"rgba(0,0,0,1)","family":"","size":17.534246575342465},"x":0,"xref":"paper"},"xaxis":{"domain":[0,1],"automargin":true,"type":"linear","autorange":false,"range":[1967.4000000000001,2024.5999999999999],"tickmode":"array","ticktext":["1970","1980","1990","2000","2010","2020"],"tickvals":[1970,1980,1990,2000,2010,2020],"categoryorder":"array","categoryarray":["1970","1980","1990","2000","2010","2020"],"nticks":null,"ticks":"outside","tickcolor":"rgba(51,51,51,1)","ticklen":3.6529680365296811,"tickwidth":0.66417600664176002,"showticklabels":true,"tickfont":{"color":"rgba(77,77,77,1)","family":"","size":11.68949771689498},"tickangle":-0,"showline":false,"linecolor":null,"linewidth":0,"showgrid":true,"gridcolor":"rgba(255,255,255,1)","gridwidth":0.66417600664176002,"zeroline":false,"anchor":"y","title":{"text":"Year","font":{"color":"rgba(0,0,0,1)","family":"","size":14.611872146118724}},"hoverformat":".2f"},"yaxis":{"domain":[0,1],"automargin":true,"type":"linear","autorange":false,"range":[0.60199999999999998,0.77800000000000002],"tickmode":"array","ticktext":["0.65","0.70","0.75"],"tickvals":[0.65000000000000013,0.70000000000000007,0.75000000000000011],"categoryorder":"array","categoryarray":["0.65","0.70","0.75"],"nticks":null,"ticks":"outside","tickcolor":"rgba(51,51,51,1)","ticklen":3.6529680365296811,"tickwidth":0.66417600664176002,"showticklabels":true,"tickfont":{"color":"rgba(77,77,77,1)","family":"","size":11.68949771689498},"tickangle":-0,"showline":false,"linecolor":null,"linewidth":0,"showgrid":true,"gridcolor":"rgba(255,255,255,1)","gridwidth":0.66417600664176002,"zeroline":false,"anchor":"x","title":{"text":"Weight per Person (lbs)","font":{"color":"rgba(0,0,0,1)","family":"","size":14.611872146118724}},"hoverformat":".2f"},"shapes":[{"type":"rect","fillcolor":null,"line":{"color":null,"width":0,"linetype":[]},"yref":"paper","xref":"paper","x0":0,"x1":1,"y0":0,"y1":1}],"showlegend":false,"legend":{"bgcolor":"rgba(255,255,255,1)","bordercolor":"transparent","borderwidth":1.8897637795275593,"font":{"color":"rgba(0,0,0,1)","family":"","size":11.68949771689498}},"hovermode":"closest","barmode":"relative"},"config":{"doubleClick":"reset","modeBarButtonsToAdd":["hoverclosest","hovercompare"],"showSendToCloud":false},"source":"A","attrs":{"3fa03e8154d8":{"x":{},"y":{},"type":"scatter"}},"cur_data":"3fa03e8154d8","visdat":{"3fa03e8154d8":["function (y) ","x"]},"highlight":{"on":"plotly_click","persistent":false,"dynamic":false,"selectize":false,"opacityDim":0.20000000000000001,"selected":{"opacity":1},"debounce":0},"shinyEvents":["plotly_hover","plotly_click","plotly_selected","plotly_relayout","plotly_brushed","plotly_brushing","plotly_clickannotation","plotly_doubleclick","plotly_deselect","plotly_afterplot","plotly_sunburstclick"],"base_url":"https://plot.ly"},"evals":[],"jsHooks":[]}</script>

Now for some python fun. Let’s begin by loading R’s python package.

``` r
# if it is not already installed, run the following command
# install.packages('reticulate', repos="http://cran.us.r-project.org")
library(reticulate)
```

Load python’s respective packages.

``` python
# if they are not installed, run the following command
# pip install pandas matplotlib
import pandas as pd
import matplotlib.pyplot as plt
```

Now, we’re going to load up a CSV containing the cost of vegetables.

``` python
df = pd.read_csv('datasets/veggies-2020.csv')
df.info()
```

    ## <class 'pandas.core.frame.DataFrame'>
    ## RangeIndex: 93 entries, 0 to 92
    ## Data columns (total 8 columns):
    ##  #   Column              Non-Null Count  Dtype  
    ## ---  ------              --------------  -----  
    ##  0   Vegetable           93 non-null     object 
    ##  1   Form                93 non-null     object 
    ##  2   RetailPrice         93 non-null     float64
    ##  3   RetailPriceUnit     93 non-null     object 
    ##  4   Yield               93 non-null     float64
    ##  5   CupEquivalentSize   93 non-null     float64
    ##  6   CupEquivalentUnit   93 non-null     object 
    ##  7   CupEquivalentPrice  93 non-null     float64
    ## dtypes: float64(4), object(4)
    ## memory usage: 5.9+ KB

Now we have different forms of vegetables, retail price, and so much more. The key columns here are going to be the name, Form, RetailPrice, and Yield. Why yield? Well, depending on how you use them, you may lose a good amount of of what you purchased. Something not usually considered when buying vegetables.

## Additional Questions

1.  Why did 2008 and 2013 have the largest percentage change?
2.  In the line chart, we see a steep drop at 1980, what happened?
3.  In the line chart, we see a spike at 2020, what happened?
