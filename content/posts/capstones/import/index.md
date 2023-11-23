---
slug: google-data-analytics
title: Food Import Cost Case Study
date: 2023-11-12
lastmod: 2023-11-21
categories: ["R"]
tags: ["R", "ggplot"]
---

<script src="{{< blogdown/postref >}}index_files/htmlwidgets/htmlwidgets.js"></script>
<script src="{{< blogdown/postref >}}index_files/plotly-binding/plotly.js"></script>
<script src="{{< blogdown/postref >}}index_files/typedarray/typedarray.min.js"></script>
<script src="{{< blogdown/postref >}}index_files/jquery/jquery.min.js"></script>
<link href="{{< blogdown/postref >}}index_files/crosstalk/css/crosstalk.min.css" rel="stylesheet" />
<script src="{{< blogdown/postref >}}index_files/crosstalk/js/crosstalk.min.js"></script>
<link href="{{< blogdown/postref >}}index_files/plotly-htmlwidgets-css/plotly-htmlwidgets.css" rel="stylesheet" />
<script src="{{< blogdown/postref >}}index_files/plotly-main/plotly-latest.min.js"></script>

I chose this topic as there is a lot that goes into the production of food. My concern was raised as I noticed that people with various disorders are struggling to get the food items they can have. By this I mean, people with Crohn’s disease, restrictive diets due to genetics (dysautonomia, EDS, ANS, etc), as well as many others. Medications can make the lack of proper foods tolerable, but it does not solve the overall issue they have.
<!--more-->
\# Week 1

## Setup

The libraries to use:

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
4)  Figure out the pattern in the data (likely going to have go back over the course of a few years)

The main question, why does food cost so much? To answer this question, we need a few bits of data to start with. Those data sets are, the current cost to produce (e.g. grow fruits) and/or manufacture (e.g. process grains) the food items. From there, we need to get a data set for the cost to consumers of these same items. Considering that I live in the United States, I’ll be working with those data sets first.

## Let’s begin

Time to start gathering the data sets. Charts will appear in the order made in the order of the data sets. If more data sets are listed than there are charts, then I have not finished going through the data.

Well, we have two charts made showing the breakdown of imported goods and the year-over-year change in cost. I’m not sure which is more concerning. The spikes in import price or the amount of food still needed to be imported. Further review will be needed to provide, not only context, but an explanation as to why we saw such a sharp increase.

# End of Week 1 Update

The data is immense. As of right now, I’m still restructuring some of the data for it to be more accessible for visualization. The Vegetables and Pulses Yearbook, while extensive (50+ tables), is structured for spreadsheet use. I still have a lot of restructuring to do, so until then… Have a happy holiday!

# Week 2

The line chart is the average of imported vegetables and pulses available by US pounds per person. It was made using ggplot2 and a custom data frame.

# Data sets

[Food Imports](https://www.ers.usda.gov/data-products/u-s-food-imports/)  
[Vegetables and Pulses Yearbook](https://www.ers.usda.gov/data-products/vegetables-and-pulses-data/vegetables-and-pulses-yearbook-tables/)  
[Farm Output](https://www.ers.usda.gov/data-products/agricultural-productivity-in-the-u-s/)  
[Dairy](https://www.ers.usda.gov/data-products/dairy-data.aspx)

# Charts

<figure>
<img src="img/2022-US-Imports.png" alt="2022 US Imports (Spreadsheet)" />
<figcaption aria-hidden="true">2022 US Imports (Spreadsheet)</figcaption>
</figure>

<figure>
<img src="img/US-Imported-Food-Price-Change.png" alt="US Imported Food YoY (Spreadsheet)" />
<figcaption aria-hidden="true">US Imported Food YoY (Spreadsheet)</figcaption>
</figure>

Of course, we need some R plots.

``` r
pc <- read_xlsx('datasets/VPPT.xlsx', sheet=3)
avg <- setNames(aggregate(pc$Value, list(pc$Year), mean), c("year", "value"))


p <- ggplot(avg, aes(year, value)) +
  geom_line() +
  labs(title="Average Vegetables and Pulses Per Person") +
  ylab('Pounds per Person') + xlab('Year')

ggplotly(p)
```

<div class="plotly html-widget html-fill-item-overflow-hidden html-fill-item" id="htmlwidget-1" style="width:672px;height:480px;"></div>
<script type="application/json" data-for="htmlwidget-1">{"x":{"data":[{"x":[1970,1971,1972,1973,1974,1975,1976,1977,1978,1979,1980,1981,1982,1983,1984,1985,1986,1987,1988,1989,1990,1991,1992,1993,1994,1995,1996,1997,1998,1999,2000,2001,2002,2003,2004,2005,2006,2007,2008,2009,2010,2011,2012,2013,2014,2015,2016,2017,2018,2019,2020,2021,2022],"y":[9.0379536647137702,9.1932120562368791,9.1450613043701097,9.0574848627487849,8.9160677979117526,8.9084510071584031,9.2028732302784899,9.1716967917406205,8.9235057519514154,9.1061743887891122,7.5095226183668569,7.3746323853892841,7.4202831249205019,7.4444239928177396,7.856779535858772,7.8237856896772868,7.7979618906711012,7.9110986531381622,7.9707328209773207,8.2662091515124754,8.2663107753110907,8.41190652339181,8.3281632926170577,8.5820366040997662,8.6716207963097443,8.5721268396998358,8.7440424075242671,8.3697622452887508,8.3118878359911879,8.397161670478356,8.3840966941988171,8.1460082943808878,8.1900662395931398,8.3128024235347429,8.3759553271910718,8.2806888051824963,8.053779325282747,8.0986649752745237,7.8623003530391902,7.894175383807851,7.9718632418239732,7.6603105877473441,7.8335731089435603,7.6058985718382681,7.7254290463971431,7.5563884790379774,7.9184809712298341,7.9897994915726676,8.0751986316257796,7.7270195104326671,8.9948069451244361,8.7050538480734136,8.7455884304787546],"text":["year: 1970<br />value: 9.037954","year: 1971<br />value: 9.193212","year: 1972<br />value: 9.145061","year: 1973<br />value: 9.057485","year: 1974<br />value: 8.916068","year: 1975<br />value: 8.908451","year: 1976<br />value: 9.202873","year: 1977<br />value: 9.171697","year: 1978<br />value: 8.923506","year: 1979<br />value: 9.106174","year: 1980<br />value: 7.509523","year: 1981<br />value: 7.374632","year: 1982<br />value: 7.420283","year: 1983<br />value: 7.444424","year: 1984<br />value: 7.856780","year: 1985<br />value: 7.823786","year: 1986<br />value: 7.797962","year: 1987<br />value: 7.911099","year: 1988<br />value: 7.970733","year: 1989<br />value: 8.266209","year: 1990<br />value: 8.266311","year: 1991<br />value: 8.411907","year: 1992<br />value: 8.328163","year: 1993<br />value: 8.582037","year: 1994<br />value: 8.671621","year: 1995<br />value: 8.572127","year: 1996<br />value: 8.744042","year: 1997<br />value: 8.369762","year: 1998<br />value: 8.311888","year: 1999<br />value: 8.397162","year: 2000<br />value: 8.384097","year: 2001<br />value: 8.146008","year: 2002<br />value: 8.190066","year: 2003<br />value: 8.312802","year: 2004<br />value: 8.375955","year: 2005<br />value: 8.280689","year: 2006<br />value: 8.053779","year: 2007<br />value: 8.098665","year: 2008<br />value: 7.862300","year: 2009<br />value: 7.894175","year: 2010<br />value: 7.971863","year: 2011<br />value: 7.660311","year: 2012<br />value: 7.833573","year: 2013<br />value: 7.605899","year: 2014<br />value: 7.725429","year: 2015<br />value: 7.556388","year: 2016<br />value: 7.918481","year: 2017<br />value: 7.989799","year: 2018<br />value: 8.075199","year: 2019<br />value: 7.727020","year: 2020<br />value: 8.994807","year: 2021<br />value: 8.705054","year: 2022<br />value: 8.745588"],"type":"scatter","mode":"lines","line":{"width":1.8897637795275593,"color":"rgba(0,0,0,1)","dash":"solid"},"hoveron":"points","showlegend":false,"xaxis":"x","yaxis":"y","hoverinfo":"text","frame":null}],"layout":{"margin":{"t":43.762557077625573,"r":7.3059360730593621,"b":40.182648401826491,"l":43.105022831050235},"plot_bgcolor":"rgba(235,235,235,1)","paper_bgcolor":"rgba(255,255,255,1)","font":{"color":"rgba(0,0,0,1)","family":"","size":14.611872146118724},"title":{"text":"Average Vegetables and Pulses Per Person","font":{"color":"rgba(0,0,0,1)","family":"","size":17.534246575342465},"x":0,"xref":"paper"},"xaxis":{"domain":[0,1],"automargin":true,"type":"linear","autorange":false,"range":[1967.4000000000001,2024.5999999999999],"tickmode":"array","ticktext":["1970","1980","1990","2000","2010","2020"],"tickvals":[1970,1980,1990,2000,2010,2020],"categoryorder":"array","categoryarray":["1970","1980","1990","2000","2010","2020"],"nticks":null,"ticks":"outside","tickcolor":"rgba(51,51,51,1)","ticklen":3.6529680365296811,"tickwidth":0.66417600664176002,"showticklabels":true,"tickfont":{"color":"rgba(77,77,77,1)","family":"","size":11.68949771689498},"tickangle":-0,"showline":false,"linecolor":null,"linewidth":0,"showgrid":true,"gridcolor":"rgba(255,255,255,1)","gridwidth":0.66417600664176002,"zeroline":false,"anchor":"y","title":{"text":"Year","font":{"color":"rgba(0,0,0,1)","family":"","size":14.611872146118724}},"hoverformat":".2f"},"yaxis":{"domain":[0,1],"automargin":true,"type":"linear","autorange":false,"range":[7.2832203431448237,9.2942852725229503],"tickmode":"array","ticktext":["7.5","8.0","8.5","9.0"],"tickvals":[7.5,8,8.5,9],"categoryorder":"array","categoryarray":["7.5","8.0","8.5","9.0"],"nticks":null,"ticks":"outside","tickcolor":"rgba(51,51,51,1)","ticklen":3.6529680365296811,"tickwidth":0.66417600664176002,"showticklabels":true,"tickfont":{"color":"rgba(77,77,77,1)","family":"","size":11.68949771689498},"tickangle":-0,"showline":false,"linecolor":null,"linewidth":0,"showgrid":true,"gridcolor":"rgba(255,255,255,1)","gridwidth":0.66417600664176002,"zeroline":false,"anchor":"x","title":{"text":"Pounds per Person","font":{"color":"rgba(0,0,0,1)","family":"","size":14.611872146118724}},"hoverformat":".2f"},"shapes":[{"type":"rect","fillcolor":null,"line":{"color":null,"width":0,"linetype":[]},"yref":"paper","xref":"paper","x0":0,"x1":1,"y0":0,"y1":1}],"showlegend":false,"legend":{"bgcolor":"rgba(255,255,255,1)","bordercolor":"transparent","borderwidth":1.8897637795275593,"font":{"color":"rgba(0,0,0,1)","family":"","size":11.68949771689498}},"hovermode":"closest","barmode":"relative"},"config":{"doubleClick":"reset","modeBarButtonsToAdd":["hoverclosest","hovercompare"],"showSendToCloud":false},"source":"A","attrs":{"88420c244c2":{"x":{},"y":{},"type":"scatter"}},"cur_data":"88420c244c2","visdat":{"88420c244c2":["function (y) ","x"]},"highlight":{"on":"plotly_click","persistent":false,"dynamic":false,"selectize":false,"opacityDim":0.20000000000000001,"selected":{"opacity":1},"debounce":0},"shinyEvents":["plotly_hover","plotly_click","plotly_selected","plotly_relayout","plotly_brushed","plotly_brushing","plotly_clickannotation","plotly_doubleclick","plotly_deselect","plotly_afterplot","plotly_sunburstclick"],"base_url":"https://plot.ly"},"evals":[],"jsHooks":[]}</script>

(interactive ggplot with custom data frame using second data set sheet 3)

## Additional Questions

1.  Why did 2008 and 2013 have the largest percentage change?
2.  In the line chart, we see a steep drop at 1980, what happened?
3.  In the line chart, we see a spike at 2020, what happened?