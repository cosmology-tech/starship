import readline from 'readline';

export interface Question {
  name: string;
  type?: string; // This can be used for further customizations like validating input based on type
}

export class Inquirerer {
  private rl: readline.Interface;

  constructor() {
    this.rl = readline.createInterface({
      input: process.stdin,
      output: process.stdout
    });
  }

  // Method to prompt for missing parameters
  public async prompt<T extends object>(params: T, questions: Question[]): Promise<T> {
    const obj: any = { ...params };

    for (const question of questions) {
      if (obj[question.name] === undefined) {
        obj[question.name] = await new Promise<string>((resolve) => {
          this.rl.question(`Enter ${question.name}: `, resolve);
        });
      }
    }

    return obj as T;
  }

  // Method to cleanly close the readline interface
  public close() {
    this.rl.close();
  }
}
