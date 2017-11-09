export interface NewAnswers {
    NewAnswers: NewAnswer[];
}

export interface NewAnswer {
    IdChoice: number;
}

export interface GivenAnswer {
    IdChoice: number;
    Choice: string;
    Amount: number;
}

export interface GivenAnswers {
    GivenAnswers: GivenAnswer[];
}
