fn matadd (a: [[i32;3];3], b: [[i32;3];3]) -> [[i32;3];3] {
    let mut retval: [[i32;3];3] = [[0;3];3];
    for col in 0..3 {
        for row in 0..3 {
            retval[col][row] = a[col][row] + b[col][row];
        }
    }
    retval

}

fn matmul (a: [[i32;3];3], b: [[i32;3];3]) -> [[i32;3];3] {
    let mut retval: [[i32;3];3] = [[0;3];3];
    for col in 0..3 {
        for row in 0..3 {
            for k in 0..3 {
                retval[col][row] += a[col][k] * b[k][row];
            }
        }
    }
    retval
}


fn main() {
    let mat1 = [[1,2,3],[4,5,6],[-1,-2,-3]];
    let mat2 = [[7,8,9],[10,11,12],[-12,-11,-10]];
    let mat = matmul (mat1, mat2);
    println!("{:?}", mat[0][0]);
    println!("{:?}", mat[1]); 
    println!("{:?}", mat); 
}
