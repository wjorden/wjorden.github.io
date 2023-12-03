---
slug: SBA Loans Post Pandemic
title: How much has PPP really hurt the economy?
date: 2023-12-02
lastmod: ["lastmod", ":git", "date", "publishDate"]
tags: ["MSSQL", "T-SQL", "SSIS", "SSMS"]
layout: single
---

This started out as a guided project to learn T-SQL(MSSQLSERVER), SQL Server Integration Service (SSIS), and SQL Server Management Studio (SSMS). However, I realized something a little... weird. Doing a little digging, we know that the Small Business Administration (SBA) is a federal program. The SBA connects entrepreneurs with lenders to help small businesses keep their doors open and employees paid. A little more digging and we find articles saying the same thing I'm going to say. While it had it's use and it helped keep the economy from collapsing, I couldn't help but think about the affect this had on the national debt.

<!--more-->

## Notes 

1. The code used in this deep dive is located in my [github](https://github.com/wjorden/sql-projects/tree/main/TSQL/sba). I'll also have some code here, however it will not be in its entirety.
2. All of the results here, are raw data, meaning no formatting.
3. All references are linked where they are highlighted. There's a lot of information here, so I didn't want to make anyone scroll to the bottom for the references.

First things first, what datasets are we working with? I worked with the newest data sets from [SBA](https://data.sba.gov/dataset/ppp-foia) and the [table data](https://www.sba.gov/document/support-table-size-standards) to help convert this to CSV and link across tables.

They are excel sheets, so you will have to convert them yourself as the converted data is to large for me to host or upload. 

So, I'll reiterate, that this was a guided project to learn TSQL. This meant converting the xlsx files to csv and using SSIS to transfer the files to a locally hosted database. From there, SSMS was used to get the next part of this.

## Train of Thought

The PPP program, backed by SBA, was used to help keep people from losing their jobs.  We know the SBA is a federal program. How does the federal government get its funds? According to the [Federal Reserve](https://fiscaldata.treasury.gov/americas-finance-guide/government-revenue/), 

> The primary sources of revenue for the U.S. government are individual and corporate taxes, and taxes that are dedicated to funding Social Security and Medicare. This revenue is used to fund a variety of goods, programs, and services to support the American public and pay interest incurred from borrowing. 

This means that taxpayer money is what backed these loans. The program ended May 31, 2021. If we look at the data:
```
SELECT 
	year(DateApproved) [Year Approved],
	count(LoanNumber) [Number Approved],
	SUM(InitialApprovalAmount) [Approved Amount],
	AVG(InitialApprovalAmount) [Average Loan]
FROM [SBA_Proj].[dbo].[sba_public_data]
WHERE year(DateApproved) = 2020  
GROUP BY year(DateApproved)

UNION

SELECT 
	year(DateApproved) [Year Approved],
	count(LoanNumber) [Number Approved],
	SUM(InitialApprovalAmount) [Approved Amount],
	AVG(InitialApprovalAmount) [Average Loan]
FROM [SBA_Proj].[dbo].[sba_public_data]
WHERE year(DateApproved) = 2021
GROUP BY year(DateApproved);
```

the results are:
<table>
  <tr>
    <th> Year Approved </th>
    <th> Number Approved </th>
    <th> Approved Amount </th> 
    <th> Average Loan </th>  
  </tr>
    <tr>
    <td style="text-align: center">2020</td>
    <td style="text-align: center">7114596</td>
    <td style="text-align: center">1658711565886.78</td>
    <td style="text-align: center">233142.059772161</td>
  </tr>
  <tr>
    <td style="text-align: center">2021</td>
    <td style="text-align: center">7259278</td>
    <td style="text-align: center">684673295934.675</td>
    <td style="text-align: center">94316.9962542659</td>
  </tr>
</table>

This means that in 2020, the approved amount totaled over \$1.65 trillion and in 2021 \$684 billion dollars. With almost a 2.5x drop, it is clear that ending the program in May was the right move. Yes, in 5 months, almost half (41%) of 2020's PPP loan approval was met. 

## What was SBA's budget?

According to [SBA](https://www.sba.gov/sites/sbagov/files/2021-06/FY2022_SBA_Congressional_Justification-508_0.pdf), \$982 billion for 2020. What about 2021? \$60 billion less at \$921 billion.

What about the forgiven amounts? Well,

```
-- 2020
SELECT
	count(LoanNumber) [Number Approved],
	SUM(CurrentApprovalAmount) [Approved Amount],
	AVG(CurrentApprovalAmount) [Average Loan],
	SUM(ForgivenessAmount) [Amount Forgiven],
	SUM(ForgivenessAmount)/SUM(CurrentApprovalAmount) * 100 [Percent Forgiven]
FROM [SBA_Proj].[dbo].[sba_public_data]
WHERE year(DateApproved) = 2020
ORDER BY 3 DESC;

-- 2021
SELECT
	count(LoanNumber) [Number Approved],
	SUM(CurrentApprovalAmount) [Approved Amount],
	AVG(CurrentApprovalAmount) [Average Loan],
	SUM(ForgivenessAmount) [Amount Forgiven],
	SUM(ForgivenessAmount)/SUM(CurrentApprovalAmount) * 100 [Percent Forgiven]
FROM [SBA_Proj].[dbo].[sba_public_data]
WHERE year(DateApproved) = 2021
ORDER BY 3 DESC;
```
The results are:
<table>
  <tr>
    <th> Number Approved </th>
    <th> Approved Amount </th> 
    <th> Average Loan </th>
    <th> Amount Forgiven </th>
    <th> Percentage Forgiven </th>
  </tr>
  <tr>
    <td style="text-align: center">7114596</td>
    <td style="text-align: center">1649867580745.87</td>
    <td style="text-align: center">231898.983546764</td>
    <td style="text-align: center">1600700562879.34</td>
    <td style="text-align: center">97.0199415734743</td>
  </tr>
  <tr>
    <td style="text-align: center">7259278</td>
    <td style="text-align: center">684490886129.031</td>
    <td style="text-align: center">94291.8684377469</td>
    <td style="text-align: center">658243247024.956</td>
    <td style="text-align: center">96.1653778544062</td>
  </tr>
</table>

\$1.6 trillion or 97% in 2020 and \$684 billion or 96% in 2021.

## What industries received the most money?

```
-- 2020
SELECT TOP 15 d.Sector,
	count(LoanNumber) [Number Approved],
	SUM(InitialApprovalAmount) [Approved Amount],
	AVG(InitialApprovalAmount) [Average Loan]
FROM [SBA_Proj].[dbo].[sba_public_data] p
INNER JOIN [SBA_Proj].[dbo].[sba_naics_sector_code_description] d
	ON LEFT(p.NAICSCode, 2) = d.[Lookup Codes]
WHERE year(DateApproved) = 2020
GROUP BY d.Sector
ORDER BY 3 DESC;

-- 2021
SELECT TOP 15 d.Sector,
	count(LoanNumber) [Number Approved],
	SUM(InitialApprovalAmount) [Approved Amount],
	AVG(InitialApprovalAmount) [Average Loan]
FROM [SBA_Proj].[dbo].[sba_public_data] p
INNER JOIN [SBA_Proj].[dbo].[sba_naics_sector_code_description] d
	ON LEFT(p.NAICSCode, 2) = d.[Lookup Codes]
WHERE year(DateApproved) = 2021
GROUP BY d.Sector
ORDER BY 3 DESC;
```
<table><caption><strong>2020</strong></caption>
  <tr>
    <th> Sector </th>
    <th> Number Approved </th> 
    <th> Approved Amount </th>
    <th> Average Loan </th>
  </tr>
  <tr>
    <td style="text-align: center">Health Care and Social Assistance</td>
    <td style="text-align: center">778503</td>
    <td style="text-align: center">216695907176.59</td>
    <td style="text-align: center">278349.482502431</td>
  </tr>
    <tr>
    <td style="text-align: center">Construction</td>
    <td style="text-align: center">748607</td>
    <td style="text-align: center">215953766036.385</td>
    <td style="text-align: center">288474.147364885</td>
  </tr>
    <tr>
    <td style="text-align: center">Professional, Scientific and Technical Services</td>
    <td style="text-align: center">923745</td>
    <td style="text-align: center">211594691191.768</td>
    <td style="text-align: center">229061.798647644</td>
  </tr>
    <tr>
    <td style="text-align: center">Manufacturing</td>
    <td style="text-align: center">442364</td>
    <td style="text-align: center">194173625700.745</td>
    <td style="text-align: center">438945.361061807</td>
  </tr>
   <tr>
    <td style="text-align: center">Accommodation and Food Services</td>
    <td style="text-align: center">555567</td>
    <td style="text-align: center">127766403941.245</td>
    <td style="text-align: center">229974.789613575</td>
  </tr>
   <tr>
    <td style="text-align: center">Retail Trade</td>
    <td style="text-align: center">617135</td>
    <td style="text-align: center">123297307022.143</td>
    <td style="text-align: center">199789.846665873</td>
  </tr>
   <tr>
    <td style="text-align: center">Wholesale Trade</td>
    <td style="text-align: center">284766</td>
    <td style="text-align: center">94399592587.3641</td>
    <td style="text-align: center">331498.818634823</td>
  </tr>
   <tr>
    <td style="text-align: center">Administrative and Support and Waste Management and Remediation Services</td>
    <td style="text-align: center">350453</td>
    <td style="text-align: center">85708954352.5702</td>
    <td style="text-align: center">244566.188198047</td>
  </tr>
   <tr>
    <td style="text-align: center">Other Services (except Public Administration)</td>
    <td style="text-align: center">688588</td>
    <td style="text-align: center">83947746011.6802</td>
    <td style="text-align: center">121912.879706995</td>
  </tr>
   <tr>
    <td style="text-align: center">Transportation and Warehousing</td>
    <td style="text-align: center">287688</td>
    <td style="text-align: center">55525287224.5315</td>
    <td style="text-align: center">193005.225190246</td>
  </tr>
   <tr>
    <td style="text-align: center">Real Estate and Rental and Leasing</td>
    <td style="text-align: center">312508</td>
    <td style="text-align: center">44629685234.7846</td>
    <td style="text-align: center">142811.336781089</td>
  </tr>
   <tr>
    <td style="text-align: center">Educational Services</td>
    <td style="text-align: center">128364</td>
    <td style="text-align: center">42051681495.5798</td>
    <td style="text-align: center">327597.157268236</td>
  </tr>
   <tr>
    <td style="text-align: center">Finance and Insurance</td>
    <td style="text-align: center">217376</td>
    <td style="text-align: center">33283369373.3834</td>
    <td style="text-align: center">153114.27836276</td>
  </tr>
   <tr>
    <td style="text-align: center">Information</td>
    <td style="text-align: center">107015</td>
    <td style="text-align: center">31259615667.6457</td>
    <td style="text-align: center">292104.991521242</td>
  </tr>
   <tr>
    <td style="text-align: center">Arts, Entertainment and Recreation</td>
    <td style="text-align: center">157917</td>
    <td style="text-align: center">23040172443.4322</td>
    <td style="text-align: center">145900.520168393</td>
  </tr>
</table>

<table><caption><strong>2021</strong></caption>
  <tr>
    <th> Sector </th>
    <th> Number Approved </th> 
    <th> Approved Amount </th>
    <th> Average Loan </th>
  </tr>
  <tr>
    <td style="text-align: center">Accommodation and Food Services</td>
    <td style="text-align: center">637791</td>
    <td style="text-align: center">121090210958.181</td>
    <td style="text-align: center">189858.764012319</td>
  </tr>
  <tr>
    <td style="text-align: center">Construction</td>
    <td style="text-align: center">658381</td>
    <td style="text-align: center">94109470863.0322</td>
    <td style="text-align: center">142940.745348107</td>
  </tr>
  <tr>
    <td style="text-align: center">Health Care and Social Assistance</td>
    <td style="text-align: center">567114</td>
    <td style="text-align: center">72666948156.8886</td>
    <td style="text-align: center">128134.639872916</td>
  </tr>
  <tr>
    <td style="text-align: center">Manufacturing</td>
    <td style="text-align: center">305766</td>
    <td style="text-align: center">71411955312.5087</td>
    <td style="text-align: center">233551.000806201</td>
  </tr>
  <tr>
    <td style="text-align: center">Professional, Scientific and Technical Services</td>
    <td style="text-align: center">725317</td>
    <td style="text-align: center">70300793071.8097</td>
    <td style="text-align: center">96924.2318487084</td>
  </tr>
  <tr>
    <td style="text-align: center">Other Services (except Public Administration)</td>
    <td style="text-align: center">1084124</td>
    <td style="text-align: center">48180744084.4233</td>
    <td style="text-align: center">44442.0971073635</td>
  </tr>
  <tr>
    <td style="text-align: center">Retail Trade</td>
    <td style="text-align: center">477016</td>
    <td style="text-align: center">31535278062.4976</td>
    <td style="text-align: center">66109.4765427105</td>
  </tr>
   <tr>
    <td style="text-align: center">Administrative and Support and Waste Management and Remediation Services</td>
    <td style="text-align: center">403164</td>
    <td style="text-align: center">29678380840.5387</td>
    <td style="text-align: center">73613.6679875651</td>
  </tr>
   <tr>
    <td style="text-align: center">Transportation and Warehousing</td>
    <td style="text-align: center">749703</td>
    <td style="text-align: center">28377287438.4442</td>
    <td style="text-align: center">37851.3723947272</td>
  </tr>
  <tr>
    <td style="text-align: center">Wholesale Trade</td>
    <td style="text-align: center">214534</td>
    <td style="text-align: center">27902795419.2556</td>
    <td style="text-align: center">130062.346384515</td>
  </tr>
  <tr>
    <td style="text-align: center">Arts, Entertainment and Recreation/td>
    <td style="text-align: center">232593</td>
    <td style="text-align: center">16890928857.1841</td>
    <td style="text-align: center">72620.1083316528</td>
  </tr>
  <tr>
    <td style="text-align: center">Real Estate and Rental and Leasing</td>
    <td style="text-align: center">269073</td>
    <td style="text-align: center">14842826209.4627</td>
    <td style="text-align: center">55162.8227635725</td>
  </tr>
  <tr>
    <td style="text-align: center">Agriculture, Forestry, Fishing and Hunting</td>
    <td style="text-align: center">539756</td>
    <td style="text-align: center">14104163630.1185</td>
    <td style="text-align: center">26130.6287102292</td>
  </tr>
  <tr>
    <td style="text-align: center">Educational Services</td>
    <td style="text-align: center">115611</td>
    <td style="text-align: center">13915323692.5375</td>
    <td style="text-align: center">120363.319169781</td>
  </tr>
   <tr>
    <td style="text-align: center">Information</td>
    <td style="text-align: center">86332</td>
    <td style="text-align: center">11368878128.4809</td>
    <td style="text-align: center">131687.880837706</td>
  </tr>
</table>

These values are as expected. 2020 was the outbreak so Health Care is expected to top the chart. 2021 Accommodation and Food Services topped the chart, also expected, as many people were still ordering in so restaurants took a huge hit.

## Results

1. In 2020, SBA was 38.6% over budget.
2. In 2021, SBA was 25% under budget.
3. Combined there was a net loss of 23.14%.
The formula for this is:
```
% over budget = (expense - budget) / budget * 100
              = ((1.6 tril + 684 bil) - (982 bil + 921 bil)) / (982 bil + 921 bil) * 100 
```

Guess where that incurred loss went? To the national debt for taxpayers to deal with.  
Guess where that forgiven money went? To the lenders.  
