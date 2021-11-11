import Node from './node';

class Trie {
    private root: Node;
    private current: Node | undefined;

    constructor() {
        this.root = new Node();
        this.current = this.root;
    }

    insert(text: string) {
        this.current = this.root;

        for (const char of text) {
            const hasChild = this.root.children.has(char);

            if (!hasChild && !this.current) {
                this.root.setChildren(char);
                this.current = this.root.children.get(char);
                continue;
            }

            if (this.current) {
                if (!this.current.children.has(char)) {
                    this.current.setChildren(char);
                }
            }

            this.current = this.current?.children.get(char);
        }

        this.current?.setStringEnd();
        this.current = this.root;
    }

    search(text: string) {
        this.current = this.root;

        for (const char of text) {
            if (!this.current?.children.has(char)) {
                this.current = this.root;
                return false;
            }

            this.current = this.current?.children.get(char);
        }

        return !!this.current?.isComplete;
    }

    remove(text: string) {
        this.current = this.root;

        for (const char of text) {
            const hasChild = this.current?.children.has(char);

            if (!hasChild) {
                this.current = this.root;
                return;
            }

            this.current = this.current?.children.get(char);
        }

        if (this.current?.isComplete) {
            this.current?.unsetStringEnd();
        }
    }
}

export default Trie;