export function getTextByService(number: number): string {
    const textMapping: { [key: number]: string } = {
        1: "家庭與育兒福利",
        2: "教育福利",
        3: "健康與退休福利",
        4: "老人與退休福利",
        5: "低收入戶與弱勢族群",
        6: "殘疾與特殊需求福利",
        7: "就業與創業福利",
        8: "社會安全與基本生活支援",
        9: "兒童及少年福利",
        10: "其他特定族群福利",
    };

    return textMapping[number] || "其他服務";
}

export function getTextByLocation(number: number): string {
    const textMapping: { [key: number]: string } = {
        1: "台北市",
        2: "新北市",
        3: "基隆市",
        4: "桃園市",
        5: "新竹市",
        6: "新竹縣",
        7: "苗栗縣",
        8: "台中市",
        9: "彰化縣",
        10: "雲林縣",
        11: "嘉義市",
        12: "嘉義縣",
        13: "台南市",
        14: "高雄市",
        15: "屏東縣",
        16: "南投縣",
        17: "宜蘭縣",
        18: "花蓮縣",
        19: "台東縣",
    };

    return textMapping[number] || "對應文字不存在";
}