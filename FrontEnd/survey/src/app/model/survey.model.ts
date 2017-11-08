export interface Survey {
    Id: number;
    Question: string;
    Choices: Choice[];
}

export interface Choice {
    Id: number;
    Choice: string;
}
