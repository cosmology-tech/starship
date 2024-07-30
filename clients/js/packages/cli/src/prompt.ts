import readline from 'readline';

export interface Question {
  name: string;
  type?: string; // This can be used for further customizations like validating input based on type
}

export class Inquirerer {
  private rl: readline.Interface | null;
  private noTty: boolean;

  constructor(noTty: boolean = false) {
    this.noTty = noTty;
    if (!noTty) {
      this.rl = readline.createInterface({
        input: process.stdin,
        output: process.stdout
      });
    } else {
      this.rl = null;
    }
  }

  // Method to prompt for missing parameters
  public async prompt<T extends object>(
    params: T,
    questions: Question[],
    usageText?: string
  ): Promise<T> {
    const obj: any = { ...params };

    if (
      usageText &&
      Object.values(params).some((value) => value === undefined) &&
      !this.noTty
    ) {
      console.log(usageText);
    }

    for (const question of questions) {
      if (obj[question.name] === undefined) {
        if (!this.noTty) {
          if (this.rl) {
            obj[question.name] = await new Promise<string>((resolve) => {
              this.rl.question(`Enter ${question.name}: `, resolve);
            });
          } else {
            throw new Error(
              'No TTY available and a readline interface is missing.'
            );
          }
        } else {
          // Optionally handle noTty cases, e.g., set defaults or throw errors
          throw new Error(`Missing required parameter: ${question.name}`);
        }
      }
    }

    return obj as T;
  }

  // Method to cleanly close the readline interface
  public close() {
    if (this.rl) {
      this.rl.close();
    }
  }
}
