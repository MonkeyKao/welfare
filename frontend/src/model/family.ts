class user {
  role: number;
  user_name: string;

  constructor(role: number, user_name: string) {
    this.role = role;
    this.user_name = user_name;
  }
}



export default class Family {
  familyName: string;
  familyId: number;
  users: Array<user>;

  constructor(familyName: string, familyId: number, users: Array<user>) {
    this.familyName = familyName;
    this.familyId = familyId
    this.users = users;
  }
}