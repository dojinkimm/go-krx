import json

import pandas as pd

dfstockcode = pd.read_html(
    "http://kind.krx.co.kr/corpgeneral/corpList.do?method=download", header=0
)[0]

stock_information = list()
for (
    i,
    (
        name,
        symbol,
        sector,
        industry,
        listing_date,
        settlement_month,
        representative,
        homepage,
        region,
    ),
) in enumerate(
    zip(
        dfstockcode.get("회사명"),
        dfstockcode.get("종목코드"),
        dfstockcode.get("업종"),
        dfstockcode.get("주요제품"),
        dfstockcode.get("상장일"),
        dfstockcode.get("결산월"),
        dfstockcode.get("대표자명"),
        dfstockcode.get("홈페이지"),
        dfstockcode.get("지역"),
    )
):
    if type(sector) == float:
        sector = "없음"

    if type(industry) == float:
        industry = "없음"

    if type(settlement_month) == float:
        settlement_month = "없음"

    if type(representative) == float:
        representative = "없음"

    if type(homepage) == float:
        homepage = "없음"

    if type(region) == float:
        region = "없음"

    symbol = str(symbol).zfill(6)
    stock_information.append(
        {
            "name": name,
            "symbol": symbol,
            "sector": sector,
            "industry": industry,
            "listing_date": listing_date,
            "settlement_month": settlement_month,
            "representative": representative,
            "homepage": homepage,
            "region": region,
        }
    )


with open("data.json", "w", encoding='utf-8') as file:
    json.dump(stock_information, file,indent=4, ensure_ascii=False)
    file.write("\n")
