class Node {
    public children: Map<string, Node>;
    public isComplete: boolean;

    constructor() {
        this.children = new Map();
        this.isComplete = false;
    }

    setChildren(value: string) {
        const emptyNode = new Node();
        this.children.set(value, emptyNode);
    }

    setStringEnd() {
        this.isComplete = true;
    }

    unsetStringEnd() {
        this.isComplete = false;
    }

}

export default Node;