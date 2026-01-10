package common

import (
	"context"
	"fmt"
	"strings"

	"github.com/Tencent/WeKnora/internal/types"
)

type I18nKey string

const (
	I18nKeyToolResultPrefix      I18nKey = "tool_result_prefix"
	I18nKeyFinalPrompt           I18nKey = "final_prompt"
	I18nKeySynthesisError        I18nKey = "synthesis_error"
	I18nKeyReflectionPrompt      I18nKey = "reflection_prompt"
	I18nKeyWebFetchSystemMessage I18nKey = "web_fetch_system_message"
	I18nKeyWebFetchUserTemplate   I18nKey = "web_fetch_user_template"
	I18nKeyTodoCreated            I18nKey = "todo_created"
	I18nKeyTodoTask               I18nKey = "todo_task"
	I18nKeyTodoSteps              I18nKey = "todo_steps"
	I18nKeyTodoProgress           I18nKey = "todo_progress"
	I18nKeyTodoTotal              I18nKey = "todo_total"
	I18nKeyTodoCompleted          I18nKey = "todo_completed"
	I18nKeyTodoInProgress         I18nKey = "todo_in_progress"
	I18nKeyTodoPending            I18nKey = "todo_pending"
	I18nKeyTodoReminder           I18nKey = "todo_reminder"
	I18nKeyTodoRemaining          I18nKey = "todo_remaining"
	I18nKeyTodoNextSteps          I18nKey = "todo_next_steps"
	I18nKeyTodoAllDone            I18nKey = "todo_all_done"
	I18nKeyTodoAllDoneNext        I18nKey = "todo_all_done_next"
	I18nKeySearchNoResults        I18nKey = "search_no_results"
	I18nKeySearchNextSteps        I18nKey = "search_next_steps"
	I18nKeySearchStats            I18nKey = "search_stats"
	I18nKeySearchDoc              I18nKey = "search_doc"
	I18nKeySearchTotalChunks      I18nKey = "search_total_chunks"
	I18nKeySearchRecalled         I18nKey = "search_recalled"
	I18nKeySearchNotRecalled      I18nKey = "search_not_recalled"
	I18nKeySearchImageDesc        I18nKey = "search_image_desc"
	I18nKeySearchImageText        I18nKey = "search_image_text"
	I18nKeyStopSessionMessage     I18nKey = "stop_session_message"

	// Core prompts from config.yaml
	I18nKeySummaryPrompt              I18nKey = "summary_prompt"
	I18nKeyRewritePromptSystem        I18nKey = "rewrite_prompt_system"
	I18nKeyRewritePromptUser          I18nKey = "rewrite_prompt_user"
	I18nKeyGenerateSessionTitlePrompt I18nKey = "generate_session_title_prompt"
	I18nKeyFallbackResponse           I18nKey = "fallback_response"
	I18nKeyFallbackPrompt             I18nKey = "fallback_prompt"
)

var i18nMessages = map[string]map[I18nKey]string{
	"zh-CN": {
		I18nKeyToolResultPrefix: "工具 %s 返回: %s",
		I18nKeyFinalPrompt: `基于上述工具调用结果，请为用户问题生成完整答案。

用户问题: %s

要求:
1. 基于实际检索到的内容回答
2. 清晰标注信息来源 (chunk_id, 文档名)
3. 结构化组织答案
4. 如信息不足，诚实说明

现在请生成最终答案:`,
		I18nKeySynthesisError:   "抱歉，我无法生成完整的答案。",
		I18nKeyReflectionPrompt: `请评估刚才调用工具 %s 的结果，并决定下一步行动。

工具返回: %s

思考:
1. 结果是否满足需求？
2. 下一步应该做什么？`,
		I18nKeyWebFetchSystemMessage: "你是一名擅长阅读网页内容的智能助手，请根据提供的网页文本回答用户需求，严禁编造未在文本中出现的信息。",
		I18nKeyWebFetchUserTemplate: `用户请求:
%s

网页内容:
%s`,
		I18nKeyTodoCreated:    "计划已创建",
		I18nKeyTodoTask:       "任务",
		I18nKeyTodoSteps:      "计划步骤",
		I18nKeyTodoProgress:   "任务进度",
		I18nKeyTodoTotal:      "总计",
		I18nKeyTodoCompleted:  "已完成",
		I18nKeyTodoInProgress: "进行中",
		I18nKeyTodoPending:    "待处理",
		I18nKeyTodoReminder:   "重要提醒",
		I18nKeyTodoRemaining:  "还有 %d 个任务未完成！\n必须完成所有任务后才能总结或得出结论。",
		I18nKeyTodoNextSteps:  "下一步操作：",
		I18nKeyTodoAllDone:    "所有任务已完成！",
		I18nKeyTodoAllDoneNext: "现在可以：\n- 综合所有任务的发现\n- 生成完整的最终答案或报告\n- 确保所有方面都已充分研究",
		I18nKeySearchNoResults: "在 %d 个知识库中未找到相关内容。",
		I18nKeySearchNextSteps: "=== ⚠️ 重要 - 下一步 ===\n- ❌ 严禁使用训练数据或通用知识回答\n- ✅ 如果启用了网络搜索：必须使用 web_search 查找信息\n- ✅ 如果禁用了网络搜索：说明“在知识库中未找到相关信息”\n- 严禁捏造或推断答案 - 仅使用检索到的内容",
		I18nKeySearchStats:       "检索统计与建议",
		I18nKeySearchDoc:         "文档",
		I18nKeySearchTotalChunks: "总 Chunk 数",
		I18nKeySearchRecalled:    "已召回",
		I18nKeySearchNotRecalled: "未召回",
		I18nKeySearchImageDesc:   "图片描述",
		I18nKeySearchImageText:   "图片文本",
		I18nKeyStopSessionMessage: "用户停止了本次对话",

		// Core prompts
		I18nKeySummaryPrompt: `你是一个专业的智能信息检索助手，名为WeKnora。你犹如专业的高级秘书，依据检索到的信息回答用户问题，不能利用任何先验知识。
当用户提出问题时，助手会基于特定的信息进行解答。助手首先在心中思考推理过程，然后向用户提供答案。
## 回答问题规则
- 仅根据检索到的信息中的事实进行回复，不得运用任何先验知识，保持回应的客观性和准确性。
- 复杂问题和答案的按Markdown分结构展示，总述部分不需要拆分
- 如果是比较简单的答案，不需要把最终答案拆分的过于细碎
- 结果中使用的图片地址必须来自于检索到的信息，不得虚构
- 检查结果中的文字和图片是否来自于检索到的信息，如果扩展了不在检索到的信息中的内容，必须进行修改，直到得到最终答案
- 如果用户问题无法回答，必须如实告知用户，并给出合理的建议。

## 输出限制
- 以Markdown图文格式输出你的最终结果
- 输出内容要保证简短且全面，条理清晰，信息明确，不重复。`,
		I18nKeyRewritePromptSystem: `你是一个专注于指代消解和省略补全的智能助手，你的任务是根据历史对话上下文，清晰识别用户问题中的代词并替换为明确的主语，同时补全省略的关键信息。

## 改写目标
请根据历史对话，对当前用户问题进行改写，目标是：
- 进行指代消解，将"它"、"这个"、"那个"、"他"、"她"、"它们"、"他们"、"她们"等代词替换为明确的主语
- 补全省略的关键信息，确保问题语义完整
- 保持问题的原始含义和表达方式不变
- 改写后必须也是一个问题
- 改写后的问题字数控制在30字以内
- 仅输出改写后的问题，不要输出任何解释，更不要尝试回答该问题，后面有其他助手回去解答此问题`,
		I18nKeyRewritePromptUser: `## 历史对话背景
{{conversation}}

## 需要改写的用户问题
{{query}}

## 改写后的问题`,
		I18nKeyGenerateSessionTitlePrompt: `根据用户的问题，生成一个简短的会话标题。

要求：
- 5-10个字
- 提取核心主题
- 只输出标题，无需解释

用户问题：`,
		I18nKeyFallbackResponse: "抱歉，我无法回答这个问题。",
		I18nKeyFallbackPrompt: `你是一个专业、友好的AI助手。请根据你的知识直接回答用户的问题。

## 回复要求
- 直接回答用户的问题
- 简洁清晰，言之有物
- 如果涉及实时数据或个人隐私信息，诚实说明无法获取
- 使用礼貌、专业的语气

## 用户的问题是:
{{query}}`,
	},
	"en-US": {
		I18nKeyToolResultPrefix: "Tool %s returned: %s",
		I18nKeyFinalPrompt: `Based on the above tool call results, please generate a complete answer for the user's question.

User Question: %s

Requirements:
1. Answer based on the actually retrieved content
2. Clearly label information sources (chunk_id, document name)
3. Organize the answer in a structured way
4. If information is insufficient, state so honestly

Now please generate the final answer:`,
		I18nKeySynthesisError:   "Sorry, I cannot generate a complete answer.",
		I18nKeyReflectionPrompt: `Please evaluate the results of the tool %s call just made and decide on the next action.

Tool returned: %s

Thinking:
1. Does the result satisfy the requirements?
2. What should be done next?`,
		I18nKeyWebFetchSystemMessage: "You are an intelligent assistant skilled at reading web content. Please answer the user's request based on the provided web text. Strictly forbid fabricating information not present in the text.",
		I18nKeyWebFetchUserTemplate: `User Request:
%s

Web Content:
%s`,
		I18nKeyTodoCreated:    "Plan Created",
		I18nKeyTodoTask:       "Task",
		I18nKeyTodoSteps:      "Plan Steps",
		I18nKeyTodoProgress:   "Task Progress",
		I18nKeyTodoTotal:      "Total",
		I18nKeyTodoCompleted:  "Completed",
		I18nKeyTodoInProgress: "In Progress",
		I18nKeyTodoPending:    "Pending",
		I18nKeyTodoReminder:   "Important Reminder",
		I18nKeyTodoRemaining:  "There are still %d task(s) remaining!\nYou must complete all tasks before generating a summary or conclusion.",
		I18nKeyTodoNextSteps:  "Next Steps:",
		I18nKeyTodoAllDone:    "All tasks completed!",
		I18nKeyTodoAllDoneNext: "You can now:\n- Synthesize findings from all tasks\n- Generate a complete final answer or report\n- Ensure all aspects have been thoroughly researched",
		I18nKeySearchNoResults: "No relevant content found in %d knowledge base(s).",
		I18nKeySearchNextSteps: "=== ⚠️ CRITICAL - Next Steps ===\n- ❌ DO NOT use training data or general knowledge to answer\n- ✅ If web_search is enabled: You MUST use web_search to find information\n- ✅ If web_search is disabled: State 'I couldn't find relevant information in the knowledge base'\n- NEVER fabricate or infer answers - ONLY use retrieved content",
		I18nKeySearchStats:       "Search Statistics & Recommendations",
		I18nKeySearchDoc:         "Document",
		I18nKeySearchTotalChunks: "Total Chunks",
		I18nKeySearchRecalled:    "Recalled",
		I18nKeySearchNotRecalled: "Not Recalled",
		I18nKeySearchImageDesc:   "Image Description",
		I18nKeySearchImageText:   "Image Text",
		I18nKeyStopSessionMessage: "The user stopped this conversation",

		// Core prompts
		I18nKeySummaryPrompt: `You are a professional intelligent information retrieval assistant named WeKnora. You are like a professional senior secretary, answering user questions based on the retrieved information, and cannot use any prior knowledge.
When a user asks a question, the assistant will provide an answer based on specific information. The assistant first thinks through the reasoning process in their mind and then provides the answer to the user.
## Question Answering Rules
- Reply only based on the facts in the retrieved information, do not use any prior knowledge, and maintain the objectivity and accuracy of the response.
- Complex questions and answers should be displayed in a structured Markdown format; the summary part does not need to be split.
- For simple answers, do not split the final answer too finely.
- The image URLs used in the results must come from the retrieved information; do not fabricate them.
- Check if the text and images in the results come from the retrieved information. If content not in the retrieved information has been expanded, it must be modified until the final answer is obtained.
- If the user's question cannot be answered, you must inform the user truthfully and provide reasonable suggestions.

## Output Restrictions
- Output your final results in Markdown format with text and images.
- Ensure the output content is concise yet comprehensive, organized, clear, and non-repetitive.`,
		I18nKeyRewritePromptSystem: `You are an intelligent assistant focused on coreference resolution and ellipsis completion. Your task is to clearly identify pronouns in user questions and replace them with explicit subjects based on the historical conversation context, while completing missing key information.

## Rewriting Goals
Please rewrite the current user question based on the historical conversation with the goal of:
- Performing coreference resolution, replacing pronouns such as "it", "this", "that", "he", "she", "they", etc., with explicit subjects.
- Completing missing key information to ensure the semantic integrity of the question.
- Keeping the original meaning and expression of the question unchanged.
- The rewritten question must also be a question.
- Keep the length of the rewritten question within 30 words.
- Output only the rewritten question, do not output any explanation, and do not try to answer the question; other assistants will answer it later.`,
		I18nKeyRewritePromptUser: `## Historical Conversation Background
{{conversation}}

## User Question to be Rewritten
{{query}}

## Rewritten Question`,
		I18nKeyGenerateSessionTitlePrompt: `Based on the user's question, generate a short conversation title.

Requirements:
- 5-10 words
- Extract the core theme
- Output only the title, no explanation needed

User Question:`,
		I18nKeyFallbackResponse: "Sorry, I cannot answer this question.",
		I18nKeyFallbackPrompt: `You are a professional and friendly AI assistant. Please answer the user's question directly based on your knowledge.

## Reply Requirements
- Answer the user's question directly.
- Be concise, clear, and meaningful.
- If real-time data or personal privacy information is involved, state honestly that it cannot be obtained.
- Use a polite and professional tone.

## User's Question is:
{{query}}`,
	},
}

// GetLocale gets the locale from context
func GetLocale(ctx context.Context) string {
	if locale, ok := ctx.Value(types.LocaleContextKey).(string); ok {
		return locale
	}
	return "zh-CN"
}

func GetI18nMsg(ctx context.Context, key I18nKey, args ...interface{}) string {
	locale := GetLocale(ctx)

	// Normalize locale (e.g., en -> en-US)
	if strings.HasPrefix(locale, "en") {
		locale = "en-US"
	} else if strings.HasPrefix(locale, "zh") {
		locale = "zh-CN"
	}

	msgs, ok := i18nMessages[locale]
	if !ok {
		// Fallback to en-US if locale not found
		msgs = i18nMessages["en-US"]
	}

	template, ok := msgs[key]
	if !ok {
		// Fallback to en-US template if key not found in current locale
		template = i18nMessages["en-US"][key]
	}

	if len(args) > 0 {
		return fmt.Sprintf(template, args...)
	}
	return template
}
