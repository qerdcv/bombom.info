import { Member } from 'components/Members/member';

import styles from 'components/Members/members.module.scss';

import { IClubMember } from 'types';
import { ReactNode } from 'react';

interface IMembersProps {
    child?: ReactNode;
    members: Array<IClubMember>
}

export const Members = ({ members }: IMembersProps) => {
    return (
        <div className={styles.members}>
            {members.map((member: IClubMember, idx: number) => {
                return <Member member={member} cnt={idx + 1} key={member.tag}/>
            })}
        </div>
    )
};
