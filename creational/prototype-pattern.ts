class UserProfile {
  name: string;
  age: number;

  constructor(name: string, age: number) {
    this.name = name;
    this.age = age;
  }

  clone(): UserProfile {
    return new UserProfile(this.name, this.age);
  }

  display(): void {
    console.log(`Name: ${this.name}, Age: ${this.age}`);
  }
}

const originalProfile = new UserProfile("Alice", 30);
const clonedProfile = originalProfile.clone();
clonedProfile.name = "Bob";
clonedProfile.age = 25; 

originalProfile.display();
clonedProfile.display();