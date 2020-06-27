<?php
// 源数组
$list = [1987, 100, 9, 250, 9, 1024, 925];
// 排序后数组
$sorted = mergeSort($list);
// 打印数组
foreach($sorted as $int) {
    printf("%d\n", $int);
}

/**
 * merge sort
 * @param array $list
 * @return array
 */
function mergeSort(array $list=[]) {
    $count = count($list);

    // 空数组无序排序
    if($count == 0) {
        return $list;
    }

    // 返回递归处理结果
    return divideMerge($list);
}

// 递归拆分数组
function divideMerge(array $list) {
    $count = count($list);

    // 递归结束条件：数组1个时，无需排序直接返回
    if($count == 1) {
        return $list;
    }

    // 递归拆分数组
    $mid = intval($count / 2);
    $leftPart = (array)array_slice($list, 0, $mid);
    $rightPart = (array)array_slice($list, $mid);
    $leftPart = divideMerge($leftPart);
    $rightPart = divideMerge($rightPart);

    // 数组合并处理
    return sortedArrayMerge($leftPart, $rightPart);
}

/**
 * merge array sorted
 * @param array $sortedLeft
 * @param array $sortedRight
 * @return array
 */
function sortedArrayMerge(array $sortedLeft, array $sortedRight) {
    $merged = [];
    while(!empty($sortedLeft) && !empty($sortedRight)) {
        if($sortedLeft[0] < $sortedRight[0]) {
            array_push($merged, array_shift($sortedLeft));
        } else {
            array_push($merged, array_shift($sortedRight));
        }
    }

    return array_merge($merged, $sortedLeft, $sortedRight);
}