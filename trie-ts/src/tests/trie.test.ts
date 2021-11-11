import Trie from '../trie/trie';

describe("test trie", () => {
    it("should found 'vitor'", () => {
        const texts = ["vitor", "john doe", "mark", "leon", "abcde", "asdf"];
        const trie = new Trie();
        texts.forEach((text) => trie.insert(text));
        expect(trie.search("vitor")).toBeTruthy();
    });

    it("should not found 'vitorr'", () => {
        const texts = ["vitor", "john doe", "mark", "leon", "abcde", "asdf"];
        const trie = new Trie();
        texts.forEach((text) => trie.insert(text));
        expect(trie.search("vitorr")).toBeFalsy();
    });

    it("should found all texts", () => {
        const shouldFound = ["vitor", "john doe", "mark", "leon", "abcde", "asdf"];
        const other = ["bbbb", "xxxx", "Cat", "wise", "rat", "mouse"];

        const trie = new Trie();
        shouldFound.forEach((text) => trie.insert(text));
        other.forEach((text) => trie.insert(text));

        const result = shouldFound.map((text) => trie.search(text));
        const otherResult = other.map((text) => trie.search(text));
        expect(result).not.toContain(false);
        expect(otherResult).not.toContain(false);
    });

    it("should not found all texts", () => {
        const shouldNotFound = ["vitor", "john doe", "mark", "leon", "abcde", "asdf"];
        const other = ["bbbb", "xxxx", "Cat", "wise", "rat", "mouse"];

        const trie = new Trie();
        other.forEach((text) => trie.insert(text));

        const result = shouldNotFound.map((text) => trie.search(text));
        expect(result).not.toContain(true);
    });

    it("should remove 'vitor'", () => {
        const texts = ["vitor", "john doe", "mark", "leon", "abcde", "asdf", "bbbb", "xxxx", "Cat", "wise", "rat", "mouse"];
        const trie = new Trie();

        texts.forEach((text) => trie.insert(text));

        trie.remove("vitor");

        expect(trie.search("vitor")).toBeFalsy();
        expect(trie.search("leon")).toBeTruthy();
    });

    it("should remove all texts", () => {
        const shouldRemove = ["vitor", "john doe", "mark", "leon", "abcde", "asdf"];
        const other = ["bbbb", "xxxx", "Cat", "wise", "rat", "mouse"];
        const trie = new Trie();

        shouldRemove.forEach((text) => trie.insert(text));
        other.forEach((text) => trie.insert(text));

        shouldRemove.forEach((text) => trie.remove(text));

        const result = shouldRemove.map((text) => trie.search(text));
        const founds = other.map((text) => trie.search(text));

        expect(result).not.toContain(true);
        expect(founds).not.toContain(false);
    });
})