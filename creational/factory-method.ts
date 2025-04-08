abstract class Animal {
  abstract makeSound(): string;
}

class Dog extends Animal {
  makeSound(): string {
    return "Woof!";
  }
}

class Cat extends Animal {
  makeSound(): string {
    return "Meow!";
  }
}

type AnimalType = "dog" | "cat";

class AnimalFactory {
  static createAnimal(type: AnimalType): Animal {
    switch (type) {
      case "dog":
        return new Dog();
      case "cat":
        return new Cat();
      default:
        throw new Error("Unknown animal type");
    }
  }
}

const dog = AnimalFactory.createAnimal("dog");
const cat = AnimalFactory.createAnimal("cat");
console.log(dog.makeSound()); 
console.log(cat.makeSound()); 
