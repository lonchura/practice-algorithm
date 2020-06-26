<?php
// 源数组
$list = [1987, 100, 9, 250, 9, 1024, 925];
// 排序后数组
$sorted = quickSort($list);
// 打印数组
foreach($sorted as $int) {
    printf("%d\n", $int);
}

/**
 * quick sort
 * @param array $list
 * @return array
 */
function quickSort(array $list=[]) {
    // 数组<2个时，无需排序直接返回
    if(count($list) < 2) {
        return $list;
    }

    // 初始化两个数组
    // 小于等于基准值数组$left
    $leftPart = [];
    // 大于基准值数组$right
    $rightPart = [];

    // 选出基准值，默认选择数组第1个数据，并移除数组的第一个元素
    $baseVal = array_shift($list);
    // 循环剩余数组
    foreach($list as $val) {
        // 将<=基准值放入数组$left
        if($val <= $baseVal) {
            array_push($leftPart, $val);
        }
        // 将>基准值放入数组$right
        if($val > $baseVal) {
            array_push($rightPart, $val);
        }
    }

    // 返回 数组拼接(quickSort($left),[基准值],quickSort($right))
    return array_merge(quickSort($leftPart), [$baseVal], quickSort($rightPart));
}