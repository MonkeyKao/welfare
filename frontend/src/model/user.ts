export default class User {
    _id: number;
    account: string;
    name: string;
    password: string;
    birthday: Date;
    female: number; 
    location: number;
    email: string;

    constructor(
        _id: number = 0,
        account: string = "",
        name: string = "",
        password: string = "",
        birthday: Date = new Date(),
        female: number = 0,
        location: number = 0,
        email: string = ""
    ) {
        this._id = _id;
        this.account = account;
        this.name = name;
        this.password = password;
        this.birthday = birthday;
        this.female = female;
        this.location = location;
        this.email = email;
    }
}