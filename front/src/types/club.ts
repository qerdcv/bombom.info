
export interface IClubMember {
    tag: string;
    name: string;
    trophies: number;
    role: string;
    nameColor: string;
}

export interface IClub {
    tag: string;
    name: string;
    description: string;
    trophies: number;
    requiredTrophies: number;
    members: Array<IClubMember>;
    type: string;
    badgeID: number;
}
