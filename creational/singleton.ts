class Logger {
  private static instance: Logger;

  private constructor() {
    if (Logger.instance) {
      throw new Error(
        "Error: Instantiation failed: Use Logger.getInstance() instead of new."
      );
    }

    Logger.instance = this;
  }

  public static getInstance(): Logger {
    if (!Logger.instance) {
      Logger.instance = new Logger();
    }
    return Logger.instance;
  }
}

const logger1 = Logger.getInstance();
const logger2 = Logger.getInstance();
console.log(logger1 === logger2);