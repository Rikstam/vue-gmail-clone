export interface EmailItem {
    id: number;
    from: string;
    subject: string;
    body: string;
    sentAt: string;
    archived: boolean;
    read: boolean;
}

export interface changeEmailEventParams {
    toggleRead?: boolean;
    save?: boolean;
    toggleArchive?: boolean;
    changeIndex?: number;
    closeModal?: boolean;
}

export interface EmailSet {
    emails: Set<EmailItem>;
}