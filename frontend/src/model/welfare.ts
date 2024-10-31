export default class Welfare {
    id: number;
    category: number[];
    city: string;
    date: string;
    priority: number;
    title: string;
    detail: string;
    detailCondition: string;
    detailDocument: string;
    url: string;

    constructor(
        id: number = 0,
        category: number[] = [],
        city: string = "",
        date: string = "",
        priority: number = 0,
        title: string = "",
        detail: string = "",
        detailCondition: string = "",
        detailDocument: string = "",
        url: string = ""
    ) {
        this.id = id;
        this.category = category;
        this.city = city;
        this.date = date;
        this.priority = priority;
        this.title = title;
        this.detail = detail;
        this.detailCondition = detailCondition;
        this.detailDocument = detailDocument;
        this.url = url;
    }
}