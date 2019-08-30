import java.util.LinkedList;
import java.util.ArrayList;

public class Codec {

    // Encodes a tree to a single string.
    public String serialize(TreeNode root) {
        if (root == null) {
            return "";
        }
        LinkedList<TreeNode> q = new LinkedList<TreeNode>();
        TreeNode node = null;
        q.add(root);
        ArrayList<String> strArr = new ArrayList<String>();

        while (q.size() != 0) {
            node = q.remove();
            if (node == null) {
                strArr.add("null");
            } else {
                strArr.add(Integer.toString(node.val));
                q.add(node.left);
                q.add(node.right);
            }
        }
        int size = strArr.size();
        int i = size-1;
        for (i = size-1; i >= 0; --i) {
            if (!strArr.get(i).equals("null")) {
                break;
            }
        }
        return String.join(",", strArr.subList(0, i+1));
    }

    // Decodes your encoded data to tree.
    public TreeNode deserialize(String data) {
        if (data.equals("") == true) {
            return null;
        }
        String[] strArr = data.split(",");
        int size = strArr.length;
        TreeNode node = null;
        TreeNode left = null;
        TreeNode right = null;

        TreeNode root = new TreeNode(Integer.parseInt(strArr[0]));
        root.left = null;
        root.right = null;

        LinkedList<TreeNode> q = new LinkedList<TreeNode>();
        q.add(root);
        int i = 1;
        while (q.size() != 0) {
            node = q.remove();
            if (i < size) {
                if (!strArr[i].equals("null")) {
                    left = new TreeNode(Integer.parseInt(strArr[i]));
                    left.left = null;
                    left.right = null;
                    q.add(left);
                    node.left = left;
                } 
                i++;
            }
            if (i < size) {
                if (!strArr[i].equals("null")) {
                    right = new TreeNode(Integer.parseInt(strArr[i]));
                    right.left = null;
                    right.right = null;
                    q.add(right);
                    node.right = right;
                } 
                i++;
            }
        }
        return root;
    }

    public static boolean testcase(String input) {
        Codec codec = new Codec();
        TreeNode tree = codec.deserialize(input);
        String output = codec.serialize(tree);
        return input.equals(output);
    }

    public static void main(String argv[]) {
        String input = "1,2,3,4,null,5,null,null,null,6,7";
        System.out.println(testcase(input));
        input = "";
        System.out.println(testcase(input));

        input = "1,2,null,3";
        System.out.println(testcase(input));
    }
}
