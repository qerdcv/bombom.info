import { ReactNode } from 'react';

import styles from 'components/Members/member.module.scss';

import { IClubMember } from 'types';

interface IMemberProps {
    child?: ReactNode;
    cnt: number;
    member: IClubMember;
}

export const Member = ({ member, cnt }: IMemberProps) => {
    return (
        <>
            <div className={styles.member}>
                <div className={styles.memberInfo}>
                    <span className={styles.memberPlace}>{ cnt }</span>
                    <div className={styles.memberNameRolled}>
                                <span
                                    title={member.tag}
                                    className={styles.memberName}
                                    style={{
                                        color: `#${member.nameColor.slice(4)}`
                                    }}
                                >
                                    {member.name}
                                </span>
                        <span className={styles.memberRole}>{member.role}</span>
                    </div>
                </div>
                <div className={styles.memberMetaInfo}>
                    <span className="curved-block"></span>
                    <span className={styles.memberTrophies}>{member.trophies}🏆</span>
                </div>
            </div>
        </>
    );
};
