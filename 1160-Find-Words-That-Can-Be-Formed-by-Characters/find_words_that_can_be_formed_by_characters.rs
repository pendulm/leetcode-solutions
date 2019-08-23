struct Solution {}

impl Solution {
    fn is_formed_by(word: &str, chars: &str) -> bool {
        let mut char_counter = [0; 26];
        for c in chars.bytes() {
            char_counter[(c - b'a') as usize] += 1;
        }

        for c in word.bytes() {
            char_counter[(c - b'a') as usize] -= 1;
            if char_counter[(c - b'a') as usize] < 0 {
                return false;
            }
        }
        return true;
    }

    pub fn count_characters(words: Vec<String>, chars: String) -> i32 {
        let mut sum_length = 0;
        for w in words.iter() {
            if Self::is_formed_by(&w, &chars) == true {
                sum_length += w.len();
            }
        }
        return sum_length as i32;
    }
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn test_case1() {
        let words = vec![
            String::from("cat"),
            String::from("bt"),
            String::from("hat"),
            String::from("tree"),
        ];
        let chars = String::from("atach");
        let count = Solution::count_characters(words, chars);
        assert_eq!(count, 6);
    }

    #[test]
    fn test_case2() {
        let words = vec![
            String::from("hello"),
            String::from("world"),
            String::from("leetcode"),
        ];
        let chars = String::from("welldonehoneyr");
        let count = Solution::count_characters(words, chars);
        assert_eq!(count, 10);
    }

}

fn main() {}
