export interface IStack {
  stack_id: number;
  stack_name: string;
  stack_description: string;
}

export interface IProject {
  p_name: string;
  p_description: string;
  p_stack: IStack;
}
