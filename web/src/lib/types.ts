export type Deck = {
    id: string;
    title: string;
    cards: [string, string][];
    createdAt: string;
    updatedAt: string;
};

export type UserCreateParams = {
    username: string;
    email: string;
    password: string;
}

export type UserCredentials = {
    email: string;
    password: string;
}