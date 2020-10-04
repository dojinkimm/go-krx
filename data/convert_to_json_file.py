import json

import pandas as pd

dfstockcode = pd.read_html(
    "http://kind.krx.co.kr/corpgeneral/corpList.do?method=download", header=0
)[0]

stock_information = list()
for (
    i,
    (
        company_name,
        code_number,
        category,
        main_product,
        listing_date,
        settlement_month,
        representative_name,
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
    if type(category) == float:
        category = "없음"

    if type(main_product) == float:
        main_product = "없음"

    if type(settlement_month) == float:
        settlement_month = "없음"

    if type(representative_name) == float:
        representative_name = "없음"

    if type(homepage) == float:
        homepage = "없음"

    if type(region) == float:
        region = "없음"

    code_number = str(code_number).zfill(6)
    stock_information.append(
        {
            "company_name": company_name,
            "code_number": code_number,
            "category": category,
            "main_product": main_product,
            "listing_date": listing_date,
            "settlement_month": settlement_month,
            "representative_name": representative_name,
            "homepage": homepage,
            "region": region,
        }
    )


with open("data.json", "w", encoding='utf-8') as json_file:
    json.dump(stock_information, json_file, ensure_ascii=False)
