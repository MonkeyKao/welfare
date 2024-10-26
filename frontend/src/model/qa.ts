export default class User {
  Id: number;
  Question: string;
  Answer: string;

  constructor(Id: number, Question: string, Answer: string) {
    this.Id = Id;
    this.Question = Question;
    this.Answer = Answer;
  }
}
