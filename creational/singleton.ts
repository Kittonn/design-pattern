class Singleton {
  private static instance: Singleton;

  private constructor() {
    if (Singleton.instance) {
      throw new Error("Use Singleton.getInstance() to get the instance of this class.");
    }

    Singleton.instance = this;
  }

  public static getInstance(): Singleton {
    if (!Singleton.instance) {
      Singleton.instance = new Singleton();
    }

    return Singleton.instance;
  }
}

const singleton1 = Singleton.getInstance();
const singleton2 = Singleton.getInstance();
console.log(singleton1 === singleton2);
