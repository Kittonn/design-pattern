interface IUser {
  name: string;
  age?: number;
  email?: string;
  address?: string;
  phone?: string;
}

class User implements IUser {
  name: string;
  age?: number;
  email?: string;
  address?: string;
  phone?: string;

  constructor(builder: UserBuilder) {
    this.name = builder.name;
    this.age = builder.age;
    this.email = builder.email;
    this.address = builder.address;
    this.phone = builder.phone;
  }

  toString(): string {
    return `User [name=${this.name}, age=${this.age}, email=${this.email}, address=${this.address}, phone=${this.phone}]`;
  }
}

class UserBuilder {
  name: string;
  age?: number;
  email?: string;
  address?: string;
  phone?: string;

  constructor(name: string) {
    this.name = name;
  }

  setAge(age: number): UserBuilder {
    this.age = age;
    return this;
  }

  setEmail(email: string): UserBuilder {
    this.email = email;
    return this;
  }

  setAddress(address: string): UserBuilder {
    this.address = address;
    return this;
  }

  setPhone(phone: string): UserBuilder {
    this.phone = phone;
    return this;
  }

  build(): User {
    return new User(this);
  }
}

const user = new UserBuilder("John Doe").setAddress("123 Main St")
console.log(user.toString());