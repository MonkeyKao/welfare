export default class User {
    _id: number;
    account: string;
    name!: string;
    password: string;
    birthday: Date;
    female: number; 
    location: number;
    email: string;

    constructor(
        _id: number,
        account: string,
        name: string,
        password: string,
        salt: string,
        birthday: Date,
        female: number,
        location: number,
        email: string
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