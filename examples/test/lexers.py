from typing import Tuple, List
from .pygments import token, lexers


def tokenize(filename: str) -> List[Tuple[str, int, int]]:
    """
    将文件内容token化，得到一个hash元组
    :param filename: 文件名称
    :return: hash列表
    """
    # 读取文件内容
    text = ""
    with open(filename, "r", encoding="utf8") as f:
        text = f.read()

    # 词法分析
    lexer = lexers.guess_lexer_for_filename(filename, text)

    # 获取token
    tokens = lexer.get_tokens(text)

    # 处理tokens并返回
    return handle_tokens(tokens)


def text_tokenize(text_content: str, suffix: str) -> List[Tuple[str, int, int]]:
    """
    将文本内容token化，得到一个hash元组
    :param text_content: 文本内容
    :param suffix: 该文本内容对应的文件后缀
    :return: hash列表
    """
    # 词法分析
    lexer = lexers.guess_lexer_for_filename(f"test{suffix}", text_content)

    # 获取token
    tokens = lexer.get_tokens(text_content)

    # 处理tokens并返回
    return handle_tokens(tokens)


def handle_tokens(tokens) -> List[Tuple[str, int, int]]:
    """
    处理tokens
    :param tokens: token列表
    :return: 结果列表
    """
    tokens = list(tokens)

    result = []
    token_length = len(tokens)
    count1 = 0  # 标记来存储原始代码文件中每个元素的对应位置
    count2 = 0  # 标记来存储清理后的代码文本中每个元素的位置

    # 这些标记用于标记原始代码文件中抄袭的内容。
    for i in range(token_length):
        if tokens[i][0] == token.Name and not i == token_length - 1 and not tokens[i + 1][1] == '(':
            result.append(('N', count1, count2))  # 所有变量名称为“N”
            count2 += 1
        elif tokens[i][0] in token.Literal.String:
            result.append(('S', count1, count2))  # 所有字符串为'S'
            count2 += 1
        elif tokens[i][0] in token.Name.Function:
            result.append(('F', count1, count2))  # 用户定义的函数名为“F”
            count2 += 1
        elif tokens[i][0] == token.Text or tokens[i][0] in token.Comment:
            pass  # 忽略了空白和注释
        else:
            result.append((tokens[i][1], count1, count2))
            # 结果元组(每个元素，例如'def'，它在原始代码文件中的位置，在清理后的代码/文本中的位置)
            count2 += len(tokens[i][1])
        count1 += len(tokens[i][1])

    return result


def hash_list_to_text(arr: List[Tuple[str, int, int]]) -> str:
    """
    将结果元组转换为字符串文本
    :param arr: tokenize处理后的结果元组
    :return: 字符串文本
    """
    return ''.join(str(x[0]) for x in arr)
