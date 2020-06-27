<?php
// 源数组
$list = [1987, 100, 9, 250, 9, 1024, 925];
// 排序后数组
$sorted = bubbleSort($list);
// 打印数组
foreach($sorted as $int) {
    printf("%d\n", $int);
}

/**
 * bubble sort
 * @param array $list
 * @return array
 */
function bubbleSort(array $list=[]) {
    // 获取数量
    $count = count($list);

    // 冒泡n-1次
    for($i=1; $i<=$count-1; $i++) {
        // 都是从第一个开始比对
        // - 首次需要比对n-1对（12 23 34 ...）
        // - 第二次需要比对n-2对（12 23 34 ...）
        // - 第i次需要比对n-i对（12 23 34 ...）
        for($j=1; $j<=$count-$i; $j++) {
            // 如果前者大于后者，就将大的往后迁移（数组索引从0开始）
            if($list[$j-1] > $list[$j]) {
                $_tmp = $list[$j];
                $list[$j] = $list[$j-1];
                $list[$j-1] = $_tmp;
            }
        }
    }

    return $list;
}