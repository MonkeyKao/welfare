export default class User {
    ID: number;
    account: string;
    name: string;
    password: string;
    birthday: string;
    female: number; 
    location: number;
    email: string;

    constructor(
        ID: number = 0,
        account: string = "",
        name: string = "",
        password: string = "",
        birthday:string = "",
        female: number = 0,
        location: number = 0,
        email: string = ""
    ) {
        this.ID = ID;
        this.account = account;
        this.name = name;
        this.password = password;
        this.birthday = birthday;
        this.female = female;
        this.location = location;
        this.email = email;
    }
}