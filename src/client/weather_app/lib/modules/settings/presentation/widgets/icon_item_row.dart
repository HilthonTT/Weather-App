import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:weather_app/common/constants/app_colors.dart';

final class IconItemDropdownRow extends StatefulWidget {
  final String title;
  final String icon;
  final List<String> options;
  final String selectedValue;
  final Function(String) onChanged;

  const IconItemDropdownRow({
    super.key,
    required this.title,
    required this.icon,
    required this.options,
    required this.selectedValue,
    required this.onChanged,
  });

  @override
  State<IconItemDropdownRow> createState() => _IconItemDropdownRowState();
}

final class _IconItemDropdownRowState extends State<IconItemDropdownRow> {
  late String _currentValue;

  @override
  void initState() {
    super.initState();
    _currentValue = widget.selectedValue;
  }

  void _onSelect(String? newValue) {
    if (newValue == null) {
      return;
    }

    setState(() {
      _currentValue = newValue;
    });

    widget.onChanged(newValue);
  }

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.symmetric(horizontal: 20),
      child: Row(
        children: [
          Image.asset(
            widget.icon,
            width: 20,
            height: 20,
            color: AppColors.gray20,
          ),
          const SizedBox(width: 15),
          Expanded(
            child: Text(
              widget.title,
              style: TextStyle(
                color: AppColors.white,
                fontSize: 14,
                fontWeight: FontWeight.w600,
              ),
            ),
          ),
          DropdownButtonHideUnderline(
            child: DropdownButton<String>(
              dropdownColor: AppColors.gray60,
              value: _currentValue,
              icon: Padding(
                padding: const EdgeInsets.only(left: 5),
                child: Image.asset(
                  "assets/icons/next.png",
                  width: 12,
                  height: 12,
                  color: AppColors.gray30,
                ),
              ),
              style: TextStyle(color: AppColors.gray30, fontSize: 12),
              onChanged: _onSelect,
              items:
                  widget.options.map((format) {
                    return DropdownMenuItem<String>(
                      value: format,
                      child: Text(format),
                    );
                  }).toList(),
            ),
          ),
        ],
      ),
    );
  }
}

final class IconItemRow extends StatelessWidget {
  final String title;
  final String icon;
  final String value;

  const IconItemRow({
    super.key,
    required this.title,
    required this.icon,
    required this.value,
  });

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.symmetric(horizontal: 20, vertical: 8),
      child: Row(
        children: [
          Image.asset(icon, width: 20, height: 20, color: AppColors.gray20),
          const SizedBox(width: 15),
          Text(
            title,
            style: TextStyle(
              color: AppColors.white,
              fontSize: 14,
              fontWeight: FontWeight.w600,
            ),
          ),
          Expanded(
            child: Text(
              value,
              textAlign: TextAlign.right,
              style: TextStyle(
                color: AppColors.gray30,
                fontSize: 12,
                fontWeight: FontWeight.w500,
              ),
            ),
          ),
          const SizedBox(width: 8),
          Image.asset(
            "assets/icons/next.png",
            width: 12,
            height: 12,
            color: AppColors.gray30,
          ),
        ],
      ),
    );
  }
}

final class IconItemSwitchRow extends StatelessWidget {
  final String title;
  final String icon;
  final bool value;
  final Function(bool) didChange;

  const IconItemSwitchRow({
    super.key,
    required this.title,
    required this.icon,
    required this.didChange,
    required this.value,
  });

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.symmetric(horizontal: 20, vertical: 8),
      child: Row(
        children: [
          Image.asset(icon, width: 20, height: 20, color: AppColors.gray20),
          const SizedBox(width: 15),
          Text(
            title,
            style: TextStyle(
              color: AppColors.white,
              fontSize: 14,
              fontWeight: FontWeight.w600,
            ),
          ),
          const Spacer(),
          const SizedBox(width: 8),
          CupertinoSwitch(value: value, onChanged: didChange),
        ],
      ),
    );
  }
}
